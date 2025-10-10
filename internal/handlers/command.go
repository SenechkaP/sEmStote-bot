package handlers

import (
	"context"

	"github.com/SenechkaP/semstore-bot/internal/keyboards"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user := update.Message.From
	chatID := update.Message.Chat.ID

	logger.Log.Infof("User: %s %s (ID: %d)", user.FirstName, user.LastName, user.ID)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        "Доброго времени суток! Чем я могу Вам помочь?",
		ReplyMarkup: keyboards.SendMainKeyboard(),
	})
}
