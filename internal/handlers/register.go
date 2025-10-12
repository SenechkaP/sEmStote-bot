package handlers

import (
	"github.com/SenechkaP/semstore-bot/internal/calculator"
	"github.com/SenechkaP/semstore-bot/internal/handlers/callbacks"
	"github.com/SenechkaP/semstore-bot/internal/handlers/commands"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var priceMatchFunc bot.MatchFunc = func(update *models.Update) bool {
	if update == nil || update.Message == nil || update.Message.Text == "" {
		return false
	}
	chatID := update.Message.Chat.ID
	_, ok := calculator.GetPending(chatID)
	return ok
}

func RegisterMessageHandlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, commands.StartHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_to_main", bot.MatchTypeExact, callbacks.HandleBackToMain)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "order", bot.MatchTypeExact, callbacks.HandleOrder)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "rate", bot.MatchTypeExact, callbacks.HandleRate)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "way_to_link", bot.MatchTypeExact, callbacks.HandleWayToLink)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_from_photo", bot.MatchTypeExact, callbacks.HandleBackFromPhoto)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "calculate", bot.MatchTypeExact, callbacks.HandleItemType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_to_item_type", bot.MatchTypeExact, callbacks.HandleItemType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "shoes", bot.MatchTypeExact, callbacks.HandleShoesType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "clothes", bot.MatchTypeExact, callbacks.HandleClothesType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "accessories", bot.MatchTypeExact, callbacks.HandleAccessoriesType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "other_item", bot.MatchTypeExact, callbacks.HandleOtherItemType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_to_category:", bot.MatchTypePrefix, callbacks.HandleBackToCategoty)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "item:", bot.MatchTypePrefix, callbacks.HandleItemSelected)
	b.RegisterHandlerMatchFunc(priceMatchFunc, callbacks.HandlePriceInput)
}
