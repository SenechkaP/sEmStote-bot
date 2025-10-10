package keyboards

import (
	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/go-telegram/bot/models"
)

func SendMainKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Заказать 📦", CallbackData: "order"},
			},
			{
				{Text: "Посчитать заказ 💸", CallbackData: "calculate"},
				{Text: "Задать вопрос", URL: constants.AdminAccount},
			},
			{
				{Text: "Актуальный курс 💹", CallbackData: "rate"},
				{Text: "Как заказать❓", URL: constants.PoizonGuide},
			},
		},
	}
}
