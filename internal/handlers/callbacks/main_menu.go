package callbacks

import (
	"context"
	"fmt"

	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/keyboards"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/SenechkaP/semstore-bot/internal/rate"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleOrder(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func HandleItemType(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func HandleRate(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID

	rateCNY, err := rate.GetRate("CNY")
	text := fmt.Sprintf(constants.RateOutput, rateCNY)

	if err != nil {
		logger.Log.Errorf("failed to fentch CNY exchange rate: %v", err)
		text = constants.RateErrorOutput
	}

	keyboards.EditMessage(ctx, b, chatID, messageID,
		text,
		keyboards.SendBackToHomeKeyboard(),
	)
}

func HandleBackToMain(ctx context.Context, b *bot.Bot, update *models.Update) {
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
