package keyboards

import (
	"fmt"

	"github.com/go-telegram/bot/models"
)

func SendBackToHomeKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "🔙 Назад в меню", CallbackData: "back_to_main"},
			},
		},
	}
}

func SendBackToCategoryKeyboard(itemCategory string) *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "🔙 Назад", CallbackData: fmt.Sprintf("back_to_category:%s", itemCategory)},
			},
		},
	}
}
