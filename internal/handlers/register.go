package handlers

import (
	"github.com/SenechkaP/semstore-bot/internal/calculator"
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
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_to_main", bot.MatchTypeExact, handleBackToMain)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "order", bot.MatchTypeExact, handleOrder)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "rate", bot.MatchTypeExact, handleRate)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "way_to_link", bot.MatchTypeExact, handleWayToLink)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_from_photo", bot.MatchTypeExact, handleBackFromPhoto)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "calculate", bot.MatchTypeExact, handleItemType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_to_item_type", bot.MatchTypeExact, handleItemType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "shoes", bot.MatchTypeExact, handleShoesType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "clothes", bot.MatchTypeExact, handleClothesType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "accessories", bot.MatchTypeExact, handlerAccessoriesType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "other_item", bot.MatchTypeExact, handlerOtherItemType)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "item:", bot.MatchTypePrefix, handleItemSelected)
	b.RegisterHandlerMatchFunc(priceMatchFunc, handlePriceInput)
}
