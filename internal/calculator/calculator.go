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
	"github.com/SenechkaP/semstore-bot/internal/rate"
)

var items map[string]Item

func SetConfig(cfg *configs.СommissionConfig) {
	items = map[string]Item{
		"sneakers":       {Name: "Кроссовки", Сommission: cfg.СommissionForSneakers},
		"other_sneakers": {Name: "Кеды", Сommission: cfg.СommissionForSneakers},
		"boots":          {Name: "Ботинки", Сommission: cfg.СommissionForSneakers},
		"heels":          {Name: "Туфли", Сommission: cfg.СommissionForSneakers},
		"slippers":       {Name: "Тапки", Сommission: cfg.СommissionForSneakers},
		"sandals":        {Name: "Сандали", Сommission: cfg.СommissionForSneakers},

		"shirts":  {Name: "Футболка/Рубашка", Сommission: cfg.CommissionForShirts},
		"hoodies": {Name: "Толстовка/Худи", Сommission: cfg.CommissionForShirts},
		"coats":   {Name: "Пуховик/Пальто", Сommission: cfg.CommissionForShirts},
		"jackets": {Name: "Жилетка/Куртка", Сommission: cfg.CommissionForShirts},
		"pants":   {Name: "Штаны", Сommission: cfg.CommissionForShirts},
		"shorts":  {Name: "Шорты", Сommission: cfg.CommissionForShirts},
		"hats":    {Name: "Шапка/Кепка", Сommission: cfg.CommissionForShirts},
		"socks":   {Name: "Носки", Сommission: cfg.CommissionForShirts},

		"glasses":   {Name: "Очки", Сommission: cfg.CommissionForShirts},
		"watches":   {Name: "Часы", Сommission: cfg.CommissionForShirts},
		"jewelry":   {Name: "Украшение", Сommission: cfg.CommissionForShirts},
		"belts":     {Name: "Ремень", Сommission: cfg.CommissionForShirts},
		"gloves":    {Name: "Перчатки", Сommission: cfg.CommissionForShirts},
		"headdress": {Name: "Головной убор", Сommission: cfg.CommissionForShirts},

		"bags":     {Name: "Рюкзак/Сумка", Сommission: cfg.CommissionForShirts},
		"continue": {Name: "Другое", Сommission: 0},
	}
}

type Item struct {
	Name       string
	Сommission int
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
	rateCNY, _ := rate.GetRate("CNY")
	total := rateCNY*float64(price) + float64(item.Сommission)
	text := fmt.Sprintf(constants.PriceOutput,
		item.Name,
		format.FormatNumberWithDots(price),
		format.FormatNumberWithDots(int(math.Ceil(total))),
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
