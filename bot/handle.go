package bot

import (
	"strings"

	"github.com/altfoxie/labrys-bot/i18n"
	"github.com/altfoxie/labrys-bot/storage"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
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
	command, args := tu.ParseCommand(message.Text)
	arg := strings.Join(args, " ")
	locale := i18n.Get(message.From.LanguageCode)

	group := strings.HasSuffix(strings.ToLower(command), "@"+strings.ToLower(b.me.Username))

	switch command {
	case "start":
		if message.Chat.Type == "private" || group {
			logrus.WithField("user", message.From.ID).Info("Message: start.")
			return b.onStart(locale, message, arg)
		}

	case "help":
		if message.Chat.Type == "private" || group {
			logrus.WithField("user", message.From.ID).Info("Message: help.")
			return b.onHelp(locale, message, arg)
		}

	case "addvoice":
		logrus.WithField("user", message.From.ID).Info("Message: add voice.")
		return b.onAddVoice(locale, message, arg)

	case "addpasta":
		logrus.WithField("user", message.From.ID).Info("Message: add copypasta.")
		return b.onAddPasta(locale, message, arg)
	}

	return nil
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
				WithSwitchPmParameter("help"),
		)
	}
}
