package callbacks

import (
	"context"

	"github.com/SenechkaP/semstore-bot/internal/constants"
	"github.com/SenechkaP/semstore-bot/internal/keyboards"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleShoesType(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func HandleClothesType(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func HandlerAccessoriesType(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func HandlerOtherItemType(ctx context.Context, b *bot.Bot, update *models.Update) {
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
