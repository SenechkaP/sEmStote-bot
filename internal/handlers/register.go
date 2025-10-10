package handlers

import "github.com/go-telegram/bot"

func RegisterMessageHandlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_to_main", bot.MatchTypeExact, handleBackToMain)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "order", bot.MatchTypeExact, handleOrder)
}
