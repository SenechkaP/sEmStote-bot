package bot

import (
	"context"

	"github.com/SenechkaP/semstore-bot/internal/handlers"
	"github.com/go-telegram/bot"
)

type TelegramBot struct {
	bot *bot.Bot
}

func New(token string) (*TelegramBot, error) {
	b, err := bot.New(token)
	if err != nil {
		return nil, err
	}

	tgBot := &TelegramBot{
		bot: b,
	}

	return tgBot, nil
}

func (t *TelegramBot) RegisterHandlers() {
	handlers.RegisterMessageHandlers(t.bot)
}

func (t *TelegramBot) Start(ctx context.Context) {
	t.bot.Start(ctx)
}
