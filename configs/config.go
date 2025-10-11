package configs

import (
	"os"
	"strconv"

	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
	CommissionConfig СommissionConfig
}

type СommissionConfig struct {
	СommissionForSneakers int
	CommissionForShirts   int
}

func LoadConfig(envPath string) *Config {
	if envPath != "" {
		if err := godotenv.Load(envPath); err != nil {
			logger.Log.Fatalf("warning: could not load env file %s: %v", envPath, err)
		}
	}

	sneakersCom, _ := strconv.Atoi(os.Getenv("COMMISSION_FOR_SNEAKERS"))
	shirtsCom, _ := strconv.Atoi(os.Getenv("COMMISSION_FOR_SHIRTS"))

	cfg := &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		CommissionConfig: СommissionConfig{
			СommissionForSneakers: sneakersCom,
			CommissionForShirts:   shirtsCom,
		},
	}

	return cfg
}
