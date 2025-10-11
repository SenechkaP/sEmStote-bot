package rate

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
)

type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

func GetRate(currency string) (float64, error) {
	allRates, err := getExchangeRate()
	if err != nil {
		return -1, err
	}

	rate := math.Round(1/allRates.Rates[currency]*100) / 100

	if currency == "CNY" {
		rate += 0.7
	}

	return rate, nil
}

func getExchangeRate() (*ExchangeRates, error) {
	responce, err := http.Get("https://www.cbr-xml-daily.ru/latest.js")
	if err != nil {
		return nil, err
	}
	defer responce.Body.Close()
	body, err := io.ReadAll(responce.Body)
	if err != nil {
		return nil, err
	}
	var rates ExchangeRates
	json.Unmarshal(body, &rates)
	return &rates, nil
}
