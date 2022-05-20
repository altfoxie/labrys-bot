package bot

import (
	"labrys-bot/internal/i18n"
	"labrys-bot/internal/storage"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

// Handler for any updates.
func (b *Bot) handle(update telego.Update) {
	var err error
	switch {
	case update.Message != nil:
		err = b.onMessage(update.Message)
	case update.InlineQuery != nil:
		err = b.onInlineQuery(update.InlineQuery)
	}
	if err != nil {
		logrus.WithError(err).Error("Error occurred while handling update.")
	}
}

// Messages handler.
func (b *Bot) onMessage(message *telego.Message) error {
	// Handle only private messages.
	if message.Chat.Type != "private" {
		return nil
	}

	command, args := tu.ParseCommand(message.Text)
	arg := strings.Join(args, " ")
	locale := i18n.Get(message.From.LanguageCode)
	switch command {
	case "start":
		return b.onStart(locale, message, arg)
	default:
		return lo.T2(b.SendMessage(
			tu.Message(tu.ID(message.Chat.ID), locale.UnknownMessage),
		)).B
	}
}

// Inline queries handler.
func (b *Bot) onInlineQuery(query *telego.InlineQuery) error {
	command, arg, _ := strings.Cut(strings.ToLower(strings.TrimSpace(query.Query)), " ")
	arg = strings.TrimSpace(arg)
	locale := i18n.Get(query.From.LanguageCode)
	switch command {
	case "v", "voice", "в", "войс", "гс":
		logrus.WithField("query", query.Query).Info("Inline query: search voice message.")
		return onDataQuery[storage.Voice, storage.VoiceMatch](b.Bot, locale, b.storage.Voices, query, arg)

	case "p", "pasta", "paste", "copypaste", "copypasta", "п", "паста", "копипаста":
		logrus.WithField("query", query.Query).Info("Inline query: search copypasta.")
		return onDataQuery[storage.Pasta, storage.PastaMatch](b.Bot, locale, b.storage.Pastas, query, arg)

	default:
		return b.AnswerInlineQuery(
			tu.InlineQuery(query.ID).
				WithSwitchPmText(locale.CommandsList).
				WithSwitchPmParameter("inline"),
		)
	}
}
