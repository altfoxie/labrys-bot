// Package bot manages the bot.
package bot

import (
	"github.com/altfoxie/labrys-bot/internal/storage"

	"github.com/mymmrac/telego"
)

// Bot represents a Telegram bot.
type Bot struct {
	*telego.Bot
	storage *storage.Storage
}

// New creates a new bot.
func New(storage *storage.Storage, token string) (*Bot, error) {
	bot, err := telego.NewBot(token)
	if err != nil {
		return nil, err
	}
	return &Bot{bot, storage}, nil
}

// Start starts the bot.
func (b *Bot) Start() error {
	updates, err := b.UpdatesViaLongPulling(nil)
	if err != nil {
		return err
	}

	for update := range updates {
		go b.handle(update)
	}

	return nil
}
