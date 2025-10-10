package handlers

import "github.com/go-telegram/bot"

func RegisterMessageHandlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_to_main", bot.MatchTypeExact, handleBackToMain)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "order", bot.MatchTypeExact, handleOrder)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "rate", bot.MatchTypeExact, handleRate)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "way_to_link", bot.MatchTypeExact, handleWayToLink)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_from_photo", bot.MatchTypeExact, handleBackFromPhoto)
}
