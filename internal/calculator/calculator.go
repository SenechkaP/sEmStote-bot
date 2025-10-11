package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type Item struct {
	Name       string
	Сommission int
}

var items = map[string]Item{
	"sneakers":       {Name: "Кроссовки", Сommission: 200},
	"other_sneakers": {Name: "Кеды", Сommission: 150},
	"boots":          {Name: "Ботинки", Сommission: 250},
	"heels":          {Name: "Туфли", Сommission: 300},
	"slippers":       {Name: "Тапки", Сommission: 300},
	"sandals":        {Name: "Сандали", Сommission: 300},

	"shirts":  {Name: "Футболка/Рубашка", Сommission: 100},
	"hoodies": {Name: "Толстовка/Худи", Сommission: 150},
	"coats":   {Name: "Пуховик/Пальто", Сommission: 400},
	"jackets": {Name: "Жилетка/Куртка", Сommission: 400},
	"pants":   {Name: "Штаны", Сommission: 400},
	"shorts":  {Name: "Шорты", Сommission: 400},
	"hats":    {Name: "Шапка/Кепка", Сommission: 400},
	"socks":   {Name: "Носки", Сommission: 400},

	"glasses":   {Name: "Очки", Сommission: 300},
	"watches":   {Name: "Часы", Сommission: 500},
	"jewelry":   {Name: "Украшение", Сommission: 300},
	"belts":     {Name: "Ремень", Сommission: 500},
	"gloves":    {Name: "Перчатки", Сommission: 300},
	"headdress": {Name: "Головной убор", Сommission: 500},

	"bags":     {Name: "Рюкзак/Сумка", Сommission: 250},
	"continue": {Name: "Другое", Сommission: 0},
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

func Compute(item Item, price int) (int, string) {
	total := price + item.Сommission
	text := fmt.Sprintf("Вы выбрали: %s\nВведённая цена: %d\nИтоговая цена: %d", item.Name, price, total)
	return total, text
}

func ParsePositiveInt(text string) (int, error) {
	val, err := strconv.Atoi(text)
	if err != nil || val < 20 {
		return 0, errors.New("некорректная сумма: введите положительное целое число")
	}
	return val, nil
}
