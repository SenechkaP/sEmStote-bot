package calculator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/SenechkaP/semstore-bot/configs"
	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/format"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/SenechkaP/semstore-bot/internal/rate"
)

var items map[string]Item

func SetConfig(cfgCom *configs.СommissionConfig, cfgShip *configs.ShippingCostConfig) {
	items = map[string]Item{
		"sneakers":       {Name: "Кроссовки", Сommission: cfgCom.CommissionForShoes, ShippingCost: cfgShip.SneakersShippingCost},
		"other_sneakers": {Name: "Кеды", Сommission: cfgCom.CommissionForShoes, ShippingCost: cfgShip.OtherSneakersShippingCost},
		"boots":          {Name: "Ботинки", Сommission: cfgCom.CommissionForShoes, ShippingCost: cfgShip.BootsShippingCost},
		"heels":          {Name: "Туфли", Сommission: cfgCom.CommissionForShoes, ShippingCost: cfgShip.HeelsShippingCost},
		"slippers":       {Name: "Тапки", Сommission: cfgCom.CommissionForShoes, ShippingCost: cfgShip.SlippersShippingCost},
		"sandals":        {Name: "Сандали", Сommission: cfgCom.CommissionForShoes, ShippingCost: cfgShip.SandalsShippingCost},

		"shirts":  {Name: "Футболка/Рубашка", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.TShirtShippingCost},
		"hoodies": {Name: "Толстовка/Худи", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.HoodieShippingCost},
		"coats":   {Name: "Пуховик/Пальто", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.CoatShippingCost},
		"jackets": {Name: "Жилетка/Куртка", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.JacketShippingCost},
		"pants":   {Name: "Штаны", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.PantsShippingCost},
		"shorts":  {Name: "Шорты", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.ShortsShippingCost},
		"hats":    {Name: "Шапка/Кепка", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.HatsShippingCost},
		"socks":   {Name: "Носки", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.SocksShippingCost},

		"glasses":   {Name: "Очки", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.GlassesShippingCost},
		"watches":   {Name: "Часы", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.WatchesShippingCost},
		"jewelry":   {Name: "Украшение", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.JewerlyShippingCost},
		"belts":     {Name: "Ремень", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.BeltShippingCost},
		"gloves":    {Name: "Перчатки", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.GlovesShippingCost},
		"headdress": {Name: "Головной убор", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.HeaddressShippingCost},

		"bags":     {Name: "Рюкзак/Сумка", Сommission: cfgCom.CommissionForOther, ShippingCost: cfgShip.BagShippingCost},
		"continue": {Name: "Другое", Сommission: 0, ShippingCost: 0},
	}
}

type Item struct {
	Name         string
	Сommission   int
	ShippingCost int
}

var (
	mu       sync.RWMutex
	sessions = map[int64]Item{}
)

var ErrItemNotFound = errors.New("item not found")

func SetPending(chatID int64, itemID string) error {
	mu.Lock()
	defer mu.Unlock()
	item, ok := items[itemID]
	if !ok {
		return ErrItemNotFound
	}
	sessions[chatID] = item
	return nil
}

func GetPending(chatID int64) (Item, bool) {
	mu.RLock()
	defer mu.RUnlock()
	it, ok := sessions[chatID]
	return it, ok
}

func ClearPending(chatID int64) {
	mu.Lock()
	defer mu.Unlock()
	delete(sessions, chatID)
}

func Compute(item Item, price int) (float64, string) {
	rateCNY_RUB, _ := rate.GetRate("CNY")
	rateEUR_RUB, _ := rate.GetRate("EUR")
	rateRUB_EUR, _ := rate.GetRubEur()

	priceInRUB := rateCNY_RUB * float64(price)
	priceInEUR := priceInRUB * rateRUB_EUR
	logger.Log.Info("price in CNY = ", price)
	logger.Log.Info("price in EUR = ", priceInEUR)

	var additionalTax float64
	var additionalText string

	if priceInEUR > 200 {
		additionalTax = (priceInEUR-200)*0.15*rateEUR_RUB + 500
		additionalText = fmt.Sprintf(constants.OverTwoHundredEurText, format.FormatNumberWithDots(int(math.Ceil(additionalTax))))
	}

	total := priceInRUB + float64(item.Сommission) + float64(item.ShippingCost) + additionalTax
	totalString := format.FormatNumberWithDots(int(math.Ceil(total)))
	totalString = totalString[:len(totalString)-2] + "90"
	text := fmt.Sprintf(constants.PriceOutput,
		item.Name,
		format.FormatNumberWithDots(price),
		totalString,
		additionalText,
	)
	return total, text
}

func ParsePositiveInt(text string) (int, error) {
	val, err := strconv.Atoi(text)
	if err != nil || val < 20 {
		return 0, errors.New("некорректная сумма: введите положительное целое число")
	}
	return val, nil
}
