package bot

import (
	"github.com/altfoxie/labrys-bot/internal/i18n"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/samber/lo"
)

// Handler for /start command.
func (b *Bot) onStart(locale *i18n.Locale, message *telego.Message, arg string) error {
	if arg == "ignore" {
		return nil
	}

	return lo.T2(b.SendMessage(
		tu.Message(
			tu.ID(message.Chat.ID),
			lo.Ternary(arg != "inline",
				locale.Greeting, locale.InlineCommandsHelp),
		),
	)).B
}
