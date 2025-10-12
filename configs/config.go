package configs

import (
	"os"
	"strconv"

	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken   string
	CommissionConfig   СommissionConfig
	DefaultRatesConfig DefaultRatesConfig
}

type СommissionConfig struct {
	СommissionForSneakers int
	CommissionForShirts   int
}

type DefaultRatesConfig struct {
	DefaultRateForCNY float64
	DefaultRateForEUR float64
}

func LoadConfig(envPath string) *Config {
	if envPath != "" {
		if err := godotenv.Load(envPath); err != nil {
			logger.Log.Fatalf(constants.LoadEnvErrorOutput, envPath, err)
		}
	}

	sneakersCom, err1 := strconv.Atoi(os.Getenv("COMMISSION_FOR_SNEAKERS"))
	shirtsCom, err2 := strconv.Atoi(os.Getenv("COMMISSION_FOR_SHIRTS"))
	defaultRateForCNY, err3 := strconv.ParseFloat(os.Getenv("CNY_DEFAULT_RATE"), 64)
	defaultRateForEUR, err4 := strconv.ParseFloat(os.Getenv("EUR_DEFAULT_RATE"), 64)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		logger.Log.Fatal(constants.EnvVariablesErrorOutput)
	}

	cfg := &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		CommissionConfig: СommissionConfig{
			СommissionForSneakers: sneakersCom,
			CommissionForShirts:   shirtsCom,
		},
		DefaultRatesConfig: DefaultRatesConfig{
			DefaultRateForCNY: defaultRateForCNY,
			DefaultRateForEUR: defaultRateForEUR,
		},
	}

	return cfg
}
