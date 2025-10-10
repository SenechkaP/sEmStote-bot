package keyboards

import (
	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/go-telegram/bot/models"
)

func SendOrderKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Написать s3mmm_7", URL: constants.AdminAccount},
				{Text: "Где взять ссылку❓", CallbackData: "category_clothes"},
			},
			{
				{Text: "🔙 Назад в меню", CallbackData: "back_to_main"},
			},
		},
	}
}
