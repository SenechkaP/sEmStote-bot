package keyboards

import (
	"github.com/go-telegram/bot/models"
)

func SendItemTypeKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "👟 Обувь", CallbackData: "shoes"},
				{Text: "👕 Одежда", CallbackData: "clothes"},
			},
			{
				{Text: "🕶 Аксессуар", CallbackData: "accessories"},
				{Text: "👜 Рюкзак/Сумка", CallbackData: "item:itemType:bags"},
			},
			{
				{Text: "Другое", CallbackData: "other_item"},
			},
			{
				{Text: "🔙 Назад в меню", CallbackData: "back_to_main"},
			},
		},
	}
}

func SendShoesTypeKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "👟 Кроссовки", CallbackData: "item:shoesType:sneakers"},
				{Text: "👟 Кеды", CallbackData: "item:shoesType:other_sneakers"},
			},
			{
				{Text: "🥾 Ботинки", CallbackData: "item:shoesType:boots"},
				{Text: "👠 Туфли", CallbackData: "item:shoesType:heels"},
			},
			{
				{Text: "🩴 Тапки", CallbackData: "item:shoesType:slippers"},
				{Text: "👡 Сандали", CallbackData: "item:shoesType:sandals"},
			},
			{
				{Text: "🔙 Вернуться назад", CallbackData: "back_to_item_type"},
			},
		},
	}
}

func SendClothesTypeKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "👕 Футболка/Рубашка", CallbackData: "item:clothesType:shirts"},
				{Text: "👘 Толстовка/Худи", CallbackData: "item:clothesType:hoodies"},
			},
			{
				{Text: "🧥 Пуховик/Пальто", CallbackData: "item:clothesType:coats"},
				{Text: "🦺 Жилетка/Куртка", CallbackData: "item:clothesType:jackets"},
			},
			{
				{Text: "👖 Штаны", CallbackData: "item:clothesType:pants"},
				{Text: "🩳 Шорты", CallbackData: "item:clothesType:shorts"},
			},
			{
				{Text: "🧢 Шапка/Кепка", CallbackData: "item:clothesType:hats"},
				{Text: "🧦 Носки", CallbackData: "item:clothesType:socks"},
			},
			{
				{Text: "🔙 Вернуться назад", CallbackData: "back_to_item_type"},
			},
		},
	}
}

func SendAccessoriesTypeKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "👓 Очки", CallbackData: "item:accessoriesType:glasses"},
				{Text: "⌚️ Часы", CallbackData: "item:accessoriesType:watches"},
			},
			{
				{Text: "💍 Украшение", CallbackData: "item:accessoriesType:jewelry"},
				{Text: "👖 Ремень", CallbackData: "item:accessoriesType:belts"},
			},
			{
				{Text: "🧤 Перчатки", CallbackData: "item:accessoriesType:gloves"},
				{Text: "🧢 Головной убор", CallbackData: "item:accessoriesType:headdress"},
			},
			{
				{Text: "🔙 Вернуться назад", CallbackData: "back_to_item_type"},
			},
		},
	}
}

func SendOtherTypeKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Продолжить", CallbackData: "item:otherType:continue"},
			},
			{
				{Text: "🔙 Вернуться назад", CallbackData: "back_to_item_type"},
			},
		},
	}
}
