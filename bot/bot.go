// Package bot manages the bot.
package bot

import (
	"github.com/altfoxie/labrys-bot/storage"

	"github.com/mymmrac/telego"
)

// Bot represents a Telegram bot.
type Bot struct {
	*telego.Bot
	me        *telego.User
	storage   *storage.Storage
	channelID int64
}

// Params is a set of parameters to create a new bot.
type Params struct {
	Token     string
	Storage   *storage.Storage
	ChannelID int64
}

// New creates a new bot.
func New(params Params) (*Bot, error) {
	bot, err := telego.NewBot(params.Token)
	if err != nil {
		return nil, err
	}

	me, err := bot.GetMe()
	if err != nil {
		return nil, err
	}

	return &Bot{bot, me, params.Storage, params.ChannelID}, nil
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
