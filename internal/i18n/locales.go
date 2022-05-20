package i18n

const defaultLang = "ru"

var locales = map[string]*Locale{
	"ru": {
		Greeting:           "Привет! Я Labrys, бот для отправки голосовых сообщений и копипаст. Напиши мне команду /help, чтобы узнать больше.",
		UnknownMessage:     "Извини, я не понимаю. Напиши мне команду /help, чтобы узнать, как я работаю.",
		CommandsList:       "Список команд",
		InlineCommandsHelp: "TODO",
		NothingFound:       "Ничего не найдено",
	},

	"ua": {
		Greeting:           "Привіт! Я Labrys, бот для відправки голосових повідомлень та копіпаст. Напиши мені команду /help, щоб дізнатися більше.",
		UnknownMessage:     "Вибач, я не знаю. Напиши мені команду /help, щоб дізнатися, як я працюю.",
		CommandsList:       "Список команд",
		InlineCommandsHelp: "TODO",
		NothingFound:       "Нічого не знайдено",
	},

	"en": {
		Greeting:           "Hi! I'm Labrys, a bot for sending voice messages and copypastas. Type /help, to learn more.",
		UnknownMessage:     "Sorry, I don't understand. Type /help, to learn how I work.",
		CommandsList:       "Commands list",
		InlineCommandsHelp: "TODO",
		NothingFound:       "Nothing found",
	},
}
