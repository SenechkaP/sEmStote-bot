package handlers

import (
	"context"

	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/keyboards"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func handleBackToMain(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	keyboards.EditMessage(ctx, b, chatID, messageID,
		constants.Greeting,
		keyboards.SendMainKeyboard(),
	)
}

func handleOrder(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	keyboards.EditMessage(ctx, b, chatID, messageID,
		constants.StepsToMakeOrder,
		keyboards.SendOrderKeyboard(),
	)
}

func handleRate(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	keyboards.EditMessage(ctx, b, chatID, messageID,
		constants.RateInfo,
		keyboards.SendRateKeyboard(),
	)
}
