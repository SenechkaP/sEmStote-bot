package rate

import (
	"math"

	"github.com/SenechkaP/semstore-bot/internal/services"
)

var exchangeServise *services.ExchangeService

func SetService(service *services.ExchangeService) {
	exchangeServise = service
}

// Returns amount of RUB for 1 unit of currency
func GetRate(currency string) (float64, error) {
	curRate, err := exchangeServise.GetRate(currency)
	rate := math.Round(1/curRate*100) / 100

	if currency == "CNY" {
		rate += 0.7
	}

	return rate, err
}

func GetRubEur() (float64, error) {
	return exchangeServise.GetRate("EUR")
}
