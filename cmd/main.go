package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/SenechkaP/semstore-bot/configs"
	"github.com/SenechkaP/semstore-bot/internal/bot"
	"github.com/SenechkaP/semstore-bot/internal/calculator"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/SenechkaP/semstore-bot/internal/rate"
	"github.com/SenechkaP/semstore-bot/internal/redis"
	"github.com/SenechkaP/semstore-bot/internal/services"
)

func main() {
	config := configs.LoadConfig(".env")

	calculator.SetConfig(&config.CommissionConfig, &config.ShippingCostConfig)

	redisClient := redis.New(config.RedisConfig.RedisAddr, config.RedisConfig.RedisPassword)
	exchangeCache := redis.NewExchangeCache(redisClient)
	exchangeService := services.NewExchangeService(exchangeCache, &config.DefaultRatesConfig)

	rate.SetService(exchangeService)

	if err := exchangeService.UpdateRates(); err != nil {
		logger.Log.Warnf("Failed to update rates at startup: %v", err)
	}

	exchangeService.StartAutoRefresh()
	defer exchangeService.StopAutoRefresh()

	tgBot, err := bot.New(config.TelegramBotToken)
	if err != nil {
		logger.Log.Fatalf("Failed to create bot: %v", err)
	}

	tgBot.RegisterHandlers()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger.Log.Info("Starting bot...")
	tgBot.Start(ctx)
	logger.Log.Info("Bot stopped")
}
