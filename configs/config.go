package configs

import (
	"os"
	"strconv"

	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/logger"
)

type Config struct {
	TelegramBotToken   string
	RedisConfig        RedisConfig
	CommissionConfig   СommissionConfig
	DefaultRatesConfig DefaultRatesConfig
	ShippingCostConfig ShippingCostConfig
}

type RedisConfig struct {
	RedisAddr     string
	RedisPassword string
}

type СommissionConfig struct {
	CommissionForShoes int
	CommissionForOther int
}

type DefaultRatesConfig struct {
	DefaultRateRUB_CNY float64
	DefaultRateRUB_EUR float64
}

type ShippingCostConfig struct {
	TShirtShippingCost        int
	HoodieShippingCost        int
	JacketShippingCost        int
	CoatShippingCost          int
	PantsShippingCost         int
	ShortsShippingCost        int
	SocksShippingCost         int
	HatsShippingCost          int
	SneakersShippingCost      int
	OtherSneakersShippingCost int
	BootsShippingCost         int
	HeelsShippingCost         int
	SlippersShippingCost      int
	SandalsShippingCost       int
	GlassesShippingCost       int
	GlovesShippingCost        int
	JewerlyShippingCost       int
	WatchesShippingCost       int
	BeltShippingCost          int
	HeaddressShippingCost     int
	BagShippingCost           int
}

func LoadConfig() *Config {
	shoesCom := getEnvInt("COMMISSION_FOR_SHOES")
	otherCom := getEnvInt("COMMISSION_FOR_OTHER")

	defaultRateRUB_CNY := getEnvFloat("RUB_CNY_DEFAULT_RATE")
	defaultRateRUB_EUR := getEnvFloat("RUB_EUR_DEFAULT_RATE")

	redisAddr := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	var shippingCostConfig ShippingCostConfig

	shippingCostConfig.TShirtShippingCost = getEnvInt("TSHIRT_SHIPPING_COST")
	shippingCostConfig.HoodieShippingCost = getEnvInt("HOODIE_SHIPPING_COST")
	shippingCostConfig.JacketShippingCost = getEnvInt("JACKET_SHIPPING_COST")
	shippingCostConfig.CoatShippingCost = getEnvInt("COAT_SHIPPING_COST")
	shippingCostConfig.PantsShippingCost = getEnvInt("PANTS_SHIPPING_COST")
	shippingCostConfig.ShortsShippingCost = getEnvInt("SHORTS_SHIPPING_COST")
	shippingCostConfig.SocksShippingCost = getEnvInt("SOCKS_SHIPPING_COST")
	shippingCostConfig.HatsShippingCost = getEnvInt("HATS_SHIPPING_COST")
	shippingCostConfig.SneakersShippingCost = getEnvInt("SNEAKERS_SHIPPING_COST")
	shippingCostConfig.OtherSneakersShippingCost = getEnvInt("OTHER_SNEAKERS_SHIPPING_COST")
	shippingCostConfig.BootsShippingCost = getEnvInt("BOOTS_SHIPPING_COST")
	shippingCostConfig.HeelsShippingCost = getEnvInt("HEELS_SHIPPING_COST")
	shippingCostConfig.SlippersShippingCost = getEnvInt("SLIPPERS_SHIPPING_COST")
	shippingCostConfig.SandalsShippingCost = getEnvInt("SANDALS_SHIPPING_COST")
	shippingCostConfig.GlassesShippingCost = getEnvInt("GLASSES_SHIPPING_COST")
	shippingCostConfig.GlovesShippingCost = getEnvInt("GLOVES_SHIPPING_COST")
	shippingCostConfig.JewerlyShippingCost = getEnvInt("JEWERLY_SHIPPING_COST")
	shippingCostConfig.WatchesShippingCost = getEnvInt("WATCHES_SHIPPING_COST")
	shippingCostConfig.BeltShippingCost = getEnvInt("BELT_SHIPPING_COST")
	shippingCostConfig.HeaddressShippingCost = getEnvInt("HEADDRESS_SHIPPING_COST")
	shippingCostConfig.BagShippingCost = getEnvInt("BAG_SHIPPING_COST")

	cfg := &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		RedisConfig: RedisConfig{
			RedisAddr:     redisAddr,
			RedisPassword: redisPassword,
		},
		CommissionConfig: СommissionConfig{
			CommissionForShoes: shoesCom,
			CommissionForOther: otherCom,
		},
		DefaultRatesConfig: DefaultRatesConfig{
			DefaultRateRUB_CNY: defaultRateRUB_CNY,
			DefaultRateRUB_EUR: defaultRateRUB_EUR,
		},
		ShippingCostConfig: shippingCostConfig,
	}

	return cfg
}

func getEnvInt(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		logger.Log.Fatalf(constants.EnvVariablesErrorOutput, key)
	}
	return val
}

func getEnvFloat(key string) float64 {
	val, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		logger.Log.Fatalf(constants.EnvVariablesErrorOutput, key)
	}
	return val
}
