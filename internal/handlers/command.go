package handlers

import (
	"context"

	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/keyboards"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update == nil || update.Message == nil {
		return
	}

	user := update.Message.From
	chatID := update.Message.Chat.ID

	logger.Log.Infof("User: %s %s (ID: %d)", user.FirstName, user.LastName, user.ID)

	if prevMsgID, ok := GetLastMenuMessage(chatID); ok && prevMsgID != 0 {
		_, _ = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    chatID,
			MessageID: prevMsgID,
		})
		ClearLastMenuMessage(chatID)
	}

	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        constants.Greeting,
		ReplyMarkup: keyboards.SendMainKeyboard(),
	})
	if err != nil {
		logger.Log.Errorf("failed to send start message: %v", err)
		return
	}

	var sentID int
	if msg != nil {
		sentID = msg.ID
	}
	if sentID != 0 {
		SetLastMenuMessage(chatID, sentID)
	}
}
