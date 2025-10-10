package keyboards

import (
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
