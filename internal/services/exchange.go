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
		logger.Log.Warnf("Failed to read rates from redis: %v. Returning default for %s", err, currency)
		if val, ok := s.defaultRates[currency]; ok {
			return val, err
		}
		return 0, fmt.Errorf("failed to get rates and no default for %s: %w", currency, err)
	}

	if ratesData == nil {
		logger.Log.Infof("No cached rates - attempting update")
		if err := s.updateRates(); err != nil {
			logger.Log.Warnf("Update failed: %v", err)
			if val, ok := s.defaultRates[currency]; ok {
				return val, err
			}
			return 0, fmt.Errorf("no rates available for %s: %w", currency, err)
		}
		ratesData, _ = s.cache.GetRates()
	}

	if s.isExpired(ratesData) {
		logger.Log.Infof("Cached rates expired (timestamp=%d). Attempting update...", ratesData.Timestamp)
		if err := s.updateRates(); err != nil {
			logger.Log.Warnf("Failed to update rates: %v. Falling back to stale/default", err)
			if ratesData != nil {
				if v, ok := ratesData.Rates[currency]; ok {
					logger.Log.Warnf("Using stale cached rate for %s: %.6f", currency, v)
					return v, nil
				}
			}
			if v, ok := s.defaultRates[currency]; ok {
				return v, nil
			}
			return 0, fmt.Errorf("no rate available for %s", currency)
		}
		ratesData, _ = s.cache.GetRates()
	}

	if ratesData != nil {
		if v, ok := ratesData.Rates[currency]; ok {
			logger.Log.Infof("Got %s rate from Redis", currency)
			return v, nil
		}
	}

	if v, ok := s.defaultRates[currency]; ok {
		return v, fmt.Errorf("currency %s not found in rates; returning default", currency)
	}
	return 0, fmt.Errorf("currency %s not found", currency)
}

func (s *ExchangeService) isExpired(ratesData *redis.ExchangeRatesData) bool {
	if ratesData == nil {
		return true
	}
	if isDayOff(ratesData.Timestamp) {
		return time.Since(ratesData.UpdatedAt) > 12*time.Hour
	}
	expiration := time.Unix(ratesData.Timestamp, 0).Add(15 * time.Hour)
	return time.Now().After(expiration)
}

func (s *ExchangeService) UpdateRates() error {
	return s.updateRates()
}

func (s *ExchangeService) StartAutoRefresh() {
	go func() {
		for {
			ratesData, _ := s.cache.GetRates()
			if ratesData == nil {
				select {
				case <-time.After(1 * time.Second):
					_ = s.updateRates()
				case <-s.stopCh:
					return
				}
				continue
			}

			if s.isExpired(ratesData) {
				_ = s.updateRates()
				select {
				case <-s.stopCh:
					return
				default:
					continue
				}
			}

			var next time.Time
			if isDayOff(ratesData.Timestamp) {
				next = ratesData.UpdatedAt.Add(12 * time.Hour)
			} else {
				next = time.Unix(ratesData.Timestamp, 0).Add(15 * time.Hour)
			}

			sleepDuration := max(time.Until(next), 0)

			logger.Log.Infof("Next exchange rates update scheduled at %s (unix=%d) — in %s", next.UTC().Format(time.RFC3339), next.Unix(), sleepDuration)

			if sleepDuration == 0 {
				_ = s.updateRates()
				select {
				case <-s.stopCh:
					return
				default:
					continue
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
	if weekday == time.Saturday || weekday == time.Sunday {
		return true
	}
	return false
}
