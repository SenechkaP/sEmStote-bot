package callbacks

import (
	"context"
	"strings"

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

func HandleAccessoriesType(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func HandleOtherItemType(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func HandleBackToCategoty(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	chatID := update.CallbackQuery.Message.Message.Chat.ID
	messageID := update.CallbackQuery.Message.Message.ID
	itemCategory := strings.Split(update.CallbackQuery.Data, ":")[1]

	var keyboard *models.InlineKeyboardMarkup
	var text string
	switch itemCategory {
	case "itemType":
		text = constants.ChooseItemType
		keyboard = keyboards.SendItemTypeKeyboard()
	case "shoesType":
		text = constants.ChooseShoesType
		keyboard = keyboards.SendShoesTypeKeyboard()
	case "clothesType":
		text = constants.ChooseClothesType
		keyboard = keyboards.SendClothesTypeKeyboard()
	case "accessoriesType":
		text = constants.ChooseAccessoriesType
		keyboard = keyboards.SendAccessoriesTypeKeyboard()
	case "otherType":
		text = constants.OtherItemTypeText
		keyboard = keyboards.SendOtherTypeKeyboard()
	}

	keyboards.EditMessage(ctx, b, chatID, messageID,
		text,
		keyboard,
	)
}
