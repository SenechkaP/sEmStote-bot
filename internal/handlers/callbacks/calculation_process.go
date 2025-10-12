package callbacks

import (
	"context"
	"strings"

	"github.com/SenechkaP/semstore-bot/internal/calculator"
	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/handlers/commands"
	"github.com/SenechkaP/semstore-bot/internal/keyboards"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleItemSelected(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := ""
	if update.CallbackQuery != nil {
		data = update.CallbackQuery.Data
	}
	if data == "" {
		return
	}
	parts := strings.SplitN(data, ":", 3)
	if len(parts) != 3 {
		return
	}
	itemCategory := parts[1]
	itemID := parts[2]

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	if err := calculator.SetPending(chatID, itemID); err != nil {
		logger.Log.Errorf("unknown item selected: %s", itemID)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Произошла ошибка: неизвестный товар.",
		})
		return
	}

	keyboards.EditMessage(ctx, b, chatID, messageID,
		constants.EnterPrice,
		keyboards.SendBackToCategoryKeyboard(itemCategory),
	)
}

func HandlePriceInput(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.From == nil {
		return
	}
	chatID := update.Message.Chat.ID
	text := strings.TrimSpace(update.Message.Text)

	item, ok := calculator.GetPending(chatID)
	if !ok {
		return
	}

	price, err := calculator.ParsePositiveInt(text)
	if err != nil {
		_, sendErr := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Неверный формат. Введите, пожалуйста, положительное целое число, большее 20",
		})
		if sendErr != nil {
			logger.Log.Errorf("failed to send parse error: %v", sendErr)
		}
		return
	}

	_, resultText := calculator.Compute(item, price)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    chatID,
		Text:      resultText,
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		logger.Log.Errorf("failed to send calc result: %v", err)
	}
	commands.StartHandler(ctx, b, update)

	calculator.ClearPending(chatID)
}
