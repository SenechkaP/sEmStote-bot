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
)

func main() {
	config := configs.LoadConfig(".env")
	calculator.SetConfig(&config.CommissionConfig, &config.ShippingCostConfig)
	rate.SetConfig(&config.DefaultRatesConfig)

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
