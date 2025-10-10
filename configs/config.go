package configs

import (
	"os"

	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
}

func LoadConfig(envPath string) *Config {
	if envPath != "" {
		if err := godotenv.Load(envPath); err != nil {
			logger.Log.Fatalf("warning: could not load env file %s: %v", envPath, err)
		}
	}

	cfg := &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}

	return cfg
}
