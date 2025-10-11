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
				{Text: "👜 Рюкзак/Сумка", CallbackData: "item:bags"},
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
				{Text: "👟 Кроссовки", CallbackData: "item:sneakers"},
				{Text: "👟 Кеды", CallbackData: "item:other_sneakers"},
			},
			{
				{Text: "🥾 Ботинки", CallbackData: "item:boots"},
				{Text: "👠 Туфли", CallbackData: "item:heels"},
			},
			{
				{Text: "🩴 Тапки", CallbackData: "item:slippers"},
				{Text: "👡 Сандали", CallbackData: "item:sandals"},
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
				{Text: "👕 Футболка/Рубашка", CallbackData: "item:shirts"},
				{Text: "👘 Толстовка/Худи", CallbackData: "item:hoodies"},
			},
			{
				{Text: "🧥 Пуховик/Пальто", CallbackData: "item:coats"},
				{Text: "🦺 Жилетка/Куртка", CallbackData: "item:jackets"},
			},
			{
				{Text: "👖 Штаны", CallbackData: "item:pants"},
				{Text: "🩳 Шорты", CallbackData: "item:shorts"},
			},
			{
				{Text: "🧢 Шапка/Кепка", CallbackData: "item:hats"},
				{Text: "🧦 Носки", CallbackData: "item:socks"},
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
				{Text: "👓 Очки", CallbackData: "item:glasses"},
				{Text: "⌚️ Часы", CallbackData: "item:watches"},
			},
			{
				{Text: "💍 Украшение", CallbackData: "item:jewelry"},
				{Text: "👖 Ремень", CallbackData: "item:belts"},
			},
			{
				{Text: "🧤 Перчатки", CallbackData: "item:gloves"},
				{Text: "🧢 Головной убор", CallbackData: "item:headdress"},
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
				{Text: "Продолжить", CallbackData: "item:continue"},
			},
			{
				{Text: "🔙 Вернуться назад", CallbackData: "back_to_item_type"},
			},
		},
	}
}
