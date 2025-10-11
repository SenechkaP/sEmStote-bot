package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/SenechkaP/semstore-bot/configs"
	"github.com/SenechkaP/semstore-bot/internal/bot"
	"github.com/SenechkaP/semstore-bot/internal/calculator"
	"github.com/SenechkaP/semstore-bot/internal/logger"
)

func main() {
	config := configs.LoadConfig(".env")
	logger.Log.Info(config)
	calculator.SetConfig(&config.CommissionConfig)

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
