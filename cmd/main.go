package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/SenechkaP/semstore-bot/configs"
	"github.com/SenechkaP/semstore-bot/internal/bot"
	"github.com/SenechkaP/semstore-bot/internal/logger"
)

func main() {
	config := configs.LoadConfig(".env")

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
