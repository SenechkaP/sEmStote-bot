package keyboards

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func EditMessage(ctx context.Context, b *bot.Bot, chatID int64, messageID int, newText string, keyboard *models.InlineKeyboardMarkup) {
	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      chatID,
		MessageID:   messageID,
		Text:        newText,
		ReplyMarkup: keyboard,
	})
}
