package handlers

import (
	"context"

	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/keyboards"
	"github.com/SenechkaP/semstore-bot/internal/logger"
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
		keyboards.SendBackToHomeKeyboard(),
	)
}

func handleWayToLink(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	path := "internal/src/images/way_to_link.jpg"

	err := keyboards.EditPhotoWithCaption(ctx, b, chatID, messageID, path, constants.WayToLink, keyboards.SendBackToOrderKeyboard())
	if err != nil {
		logger.Log.Errorf("image %s have not been uploaded\n", path)
	}
}

func handleBackFromPhoto(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	_, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    chatID,
		MessageID: messageID,
	})
	if err != nil {
		logger.Log.Errorf("failed to delete image message: %v", err)
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        constants.StepsToMakeOrder,
		ReplyMarkup: keyboards.SendOrderKeyboard(),
	})
	if err != nil {
		logger.Log.Errorf("failed to send order menu: %v", err)
	}
}
