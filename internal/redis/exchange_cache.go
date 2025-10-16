package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/redis/go-redis/v9"
)

const ExchangeRatesKey = "exchange:rates"
const FailuresCountKey = "exchange:failures"

type ExchangeRatesData struct {
	Timestamp int64              `json:"timestamp"`
	Rates     map[string]float64 `json:"rates"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type ExchangeCache struct {
	client RedisClient
	ctx    context.Context
}

func NewExchangeCache(client RedisClient) *ExchangeCache {
	return &ExchangeCache{
		client: client,
		ctx:    context.Background(),
	}
}

func (ec *ExchangeCache) SaveRates(data *ExchangeRatesData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal rates: %w", err)
	}

	if err := ec.client.Set(ec.ctx, ExchangeRatesKey, jsonData, 0).Err(); err != nil {
		return fmt.Errorf("failed to save rates: %w", err)
	}

	logger.Log.Infof("Saved exchange rates (timestamp=%d, currencies=%d)", data.Timestamp, len(data.Rates))
	return nil
}

func (ec *ExchangeCache) GetRates() (*ExchangeRatesData, error) {
	dataBytes, err := ec.client.Get(ec.ctx, ExchangeRatesKey).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get rates: %w", err)
	}

	var data ExchangeRatesData
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal rates: %w", err)
	}

	return &data, nil
}

func (ec *ExchangeCache) RecordFailure() (int, time.Duration, error) {
	failures, err := ec.client.Incr(ec.ctx, FailuresCountKey).Result()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to record failure: %w", err)
	}

	delaySeconds := min(failures, 300)
	delay := time.Duration(delaySeconds) * time.Second
	ec.client.Expire(ec.ctx, FailuresCountKey, time.Hour)

	logger.Log.Infof("API failure recorded (attempt %d), next retry in %v", failures, delay)
	return int(failures), delay, nil
}

func (ec *ExchangeCache) ResetFailures() error {
	if err := ec.client.Del(ec.ctx, FailuresCountKey).Err(); err != nil {
		return fmt.Errorf("failed to reset failures count: %w", err)
	}
	return nil
}
