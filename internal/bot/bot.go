package bot

import (
	"context"

	"github.com/SenechkaP/semstore-bot/internal/handlers"
	"github.com/SenechkaP/semstore-bot/internal/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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
	_, err := t.bot.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{
				Command:     "start",
				Description: "Перезапустить бота",
			},
		},
	})
	if err != nil {
		logger.Log.Errorf("failed to set commands: %v", err)
	}
	t.bot.Start(ctx)
}
