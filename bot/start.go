package bot

import (
	"fmt"

	"github.com/altfoxie/labrys-bot/i18n"

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
			lo.Ternary(arg != "help",
				locale.Greeting, fmt.Sprintf(locale.Help, b.me.Username)),
		).WithParseMode(telego.ModeHTML),
	)).B
}

// Handler for /help command.
func (b *Bot) onHelp(locale *i18n.Locale, message *telego.Message, arg string) error {
	return lo.T2(b.SendMessage(
		tu.Message(
			tu.ID(message.Chat.ID),
			fmt.Sprintf(locale.Help, b.me.Username),
		).WithParseMode(telego.ModeHTML),
	)).B
}
