package keyboards

import (
	"bytes"
	"context"
	"os"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func EditMessage(ctx context.Context, b *bot.Bot, chatID int64, messageID int, newText string, keyboard *models.InlineKeyboardMarkup) {
	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      chatID,
		MessageID:   messageID,
		Text:        newText,
		ReplyMarkup: keyboard,
		ParseMode:   models.ParseModeHTML,
	})
}

func EditPhotoWithCaption(ctx context.Context, b *bot.Bot, chatID int64, messageID int, filePath, caption string, keyboard *models.InlineKeyboardMarkup) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	fileName := getFileNameFromPath(filePath)

	media := &models.InputMediaPhoto{
		Media:           "attach://" + fileName,
		Caption:         caption,
		MediaAttachment: bytes.NewReader(fileData),
	}

	params := &bot.EditMessageMediaParams{
		ChatID:      chatID,
		MessageID:   messageID,
		Media:       media,
		ReplyMarkup: keyboard,
	}

	_, err = b.EditMessageMedia(ctx, params)
	return err
}

func getFileNameFromPath(path string) string {
	pathParts := strings.Split(path, "/")
	return pathParts[len(pathParts)-1]
}
