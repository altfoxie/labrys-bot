package bot

import (
	"fmt"
	"html"

	"github.com/altfoxie/labrys-bot/i18n"
	"github.com/altfoxie/labrys-bot/storage"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/samber/lo"
)

// Handler for /addpasta command.
func (b *Bot) onAddPasta(locale *i18n.Locale, message *telego.Message, arg string) error {
	if arg == "" {
		return lo.T2(b.SendMessage(
			tu.Message(
				tu.ID(message.Chat.ID),
				locale.NoName,
			),
		)).B
	}

	text := ""
	if reply := message.ReplyToMessage; reply != nil {
		text, _ = lo.Coalesce(reply.Text, reply.Caption)
	}
	if text == "" {
		return lo.T2(b.SendMessage(
			tu.Message(
				tu.ID(message.Chat.ID),
				locale.NoReply,
			),
		)).B
	}

	var sentID int
	if b.channelID != 0 {
		if sent, err := b.SendMessage(tu.Message(
			tu.ID(b.channelID),
			fmt.Sprintf("<b>%s</b>\n\n%s", arg, html.EscapeString(text)),
		).WithParseMode(telego.ModeHTML)); err == nil {
			sentID = sent.MessageID
		}
	}

	if err := b.storage.Pastas.Create(storage.Pasta{
		Name:      arg,
		Content:   text,
		MessageID: sentID,
	}); err != nil {
		return err
	}

	return lo.T2(b.SendMessage(
		tu.Message(
			tu.ID(message.Chat.ID),
			locale.PastaAdded,
		),
	)).B
}
