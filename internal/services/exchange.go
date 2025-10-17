package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/SenechkaP/semstore-bot/configs"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/SenechkaP/semstore-bot/internal/redis"
)

type ExchangeService struct {
	cache        *redis.ExchangeCache
	defaultRates map[string]float64
	stopCh       chan struct{}
}

type cbrResponse struct {
	Timestamp int64              `json:"timestamp"`
	Rates     map[string]float64 `json:"rates"`
}

func NewExchangeService(cache *redis.ExchangeCache, cfg *configs.DefaultRatesConfig) *ExchangeService {
	return &ExchangeService{
		cache: cache,
		defaultRates: map[string]float64{
			"CNY": cfg.DefaultRateRUB_CNY,
			"EUR": cfg.DefaultRateRUB_EUR,
		},
		stopCh: make(chan struct{}),
	}
}

func (s *ExchangeService) GetRate(currency string) (float64, error) {
	ratesData, err := s.cache.GetRates()
	if err != nil {
		logger.Log.Warnf("Failed to get rates from Redis: %v", err)
		logger.Log.Warnf("Using default rate for %s: %.2f", currency, s.defaultRates[currency])
		return s.defaultRates[currency], err
	}

	now := time.Now().Unix()
	if !isDayOff(ratesData.Timestamp) && (ratesData == nil || now > ratesData.Timestamp+5) {
		logger.Log.Infof("Rates missing or expired. Updating now...")
		if err := s.updateRates(); err != nil {
			logger.Log.Warnf("Failed to update rates, fallback: %v", err)
			if ratesData != nil {
				if val, ok := ratesData.Rates[currency]; ok {
					logger.Log.Warnf("Using stale cached rate for %s: %.2f", currency, val)
					return val, nil
				}
			}
			if val, ok := s.defaultRates[currency]; ok {
				return val, nil
			}
			return 0, fmt.Errorf("no rate available for %s", currency)
		}
		ratesData, _ = s.cache.GetRates()
	}

	if val, ok := ratesData.Rates[currency]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("currency %s not found", currency)
}

func (s *ExchangeService) UpdateRates() error {
	return s.updateRates()
}

func (s *ExchangeService) StartAutoRefresh() {
	go func() {
		for {
			ratesData, _ := s.cache.GetRates()
			var sleepDuration time.Duration
			if ratesData == nil {
				sleepDuration = 1 * time.Second
			} else {
				if isDayOff(ratesData.Timestamp) {
					sleepDuration = time.Hour * 12
				} else {
					nextUpdate := time.Unix(ratesData.Timestamp+5, 0)
					sleepDuration = max(time.Until(nextUpdate), 0)
				}
			}

			select {
			case <-time.After(sleepDuration):
				_ = s.updateRates()
			case <-s.stopCh:
				return
			}
		}
	}()
}

func (s *ExchangeService) StopAutoRefresh() {
	close(s.stopCh)
}

func (s *ExchangeService) updateRates() error {
	body, err := s.getRatesFromAPI()
	if err != nil {
		_, delay, _ := s.cache.RecordFailure()
		time.Sleep(delay)
		body, err = s.getRatesFromAPI()
		if err != nil {
			return fmt.Errorf("failed 2 times to fetch API: %w", err)
		}
	}
	s.cache.ResetFailures()

	ratesData := &redis.ExchangeRatesData{
		Timestamp: body.Timestamp,
		Rates:     body.Rates,
		UpdatedAt: time.Now(),
	}
	return s.cache.SaveRates(ratesData)
}

func (s *ExchangeService) getRatesFromAPI() (*cbrResponse, error) {
	resp, err := http.Get("https://www.cbr-xml-daily.ru/latest.js")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var raw cbrResponse
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	return &raw, nil
}

func isDayOff(timestamp int64) bool {
	weekday := time.Unix(timestamp, 0).Weekday()
	if weekday == time.Friday || weekday == time.Saturday || weekday == time.Sunday {
		return true
	}
	return false
}
