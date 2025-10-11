package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/SenechkaP/semstore-bot/internal/calculator"
	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/keyboards"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/SenechkaP/semstore-bot/internal/rate"
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

	rateCNY, _ := rate.GetRate("CNY")
	text := fmt.Sprintf("Курс на сегодня: %.2f₽ за 1¥", rateCNY)

	keyboards.EditMessage(ctx, b, chatID, messageID,
		text,
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

	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        constants.StepsToMakeOrder,
		ReplyMarkup: keyboards.SendOrderKeyboard(),
	})
	if err != nil {
		logger.Log.Errorf("failed to send order menu: %v", err)
	}

	var sentID int
	if msg != nil {
		sentID = msg.ID
	}
	if sentID != 0 {
		SetLastMenuMessage(chatID, sentID)
	}
}

func handleItemType(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	keyboards.EditMessage(ctx, b, chatID, messageID,
		constants.ChooseItemType,
		keyboards.SendItemTypeKeyboard(),
	)
}

func handleShoesType(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	keyboards.EditMessage(ctx, b, chatID, messageID,
		constants.ChooseShoesType,
		keyboards.SendShoesTypeKeyboard(),
	)
}

func handleClothesType(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	keyboards.EditMessage(ctx, b, chatID, messageID,
		constants.ChooseClothesType,
		keyboards.SendClothesTypeKeyboard(),
	)
}

func handlerAccessoriesType(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	keyboards.EditMessage(ctx, b, chatID, messageID,
		constants.ChooseAccessoriesType,
		keyboards.SendAccessoriesTypeKeyboard(),
	)
}

func handlerOtherItemType(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	keyboards.EditMessage(ctx, b, chatID, messageID,
		constants.OtherItemTypeText,
		keyboards.SendOtherTypeKeyboard(),
	)
}

func handleItemSelected(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := ""
	if update.CallbackQuery != nil {
		data = update.CallbackQuery.Data
	}
	if data == "" {
		return
	}
	parts := strings.SplitN(data, ":", 2)
	if len(parts) != 2 {
		return
	}
	itemID := parts[1]

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
		&models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{},
		},
		// keyboards.SendBackToHomeKeyboard(),
	)
}

func handlePriceInput(ctx context.Context, b *bot.Bot, update *models.Update) {
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
		ChatID: chatID,
		Text:   resultText,
		// ReplyMarkup: keyboards.SendBackToHomeKeyboard(),
	})
	if err != nil {
		logger.Log.Errorf("failed to send calc result: %v", err)
	}
	startHandler(ctx, b, update)

	calculator.ClearPending(chatID)
}
