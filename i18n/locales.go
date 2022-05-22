package i18n

const defaultLang = "ru"

var locales = map[string]*Locale{
	"ru": {
		Greeting: "Привет! Я Labrys, бот для отправки голосовых сообщений и копипаст. Напиши мне команду /help, чтобы узнать больше.",
		Help: `Команды:
<code>/help</code> - повторно вызвать это сообщение.
<code>/addvoice &lt;название&gt;</code> (в ответ на аудио/файл) - добавить голосовое сообщение.
<code>/addpasta &lt;название&gt;</code> (в ответ на текст) - добавить копипасту.

Инлайн-команды:
<code>@%[1]s [v|voice|в|войс|гс] &lt;название&gt;</code> - поиск голосового сообщения.
<code>@%[1]s [p|pasta|п|паста] &lt;название&gt;</code> - поиск копипасты.`,

		NoName:     "После команды нужно написать название.",
		NoReply:    "Эту команду можно использовать только в ответ на сообщение.",
		VoiceAdded: "Голосовое сообщение добавлено.",
		PastaAdded: "Копипаста добавлена.",

		CommandsList: "Список команд",
		NothingFound: "Ничего не найдено",
	},

	"ua": {
		Greeting: "Привіт! Я Labrys, бот для відправки голосових повідомлень та копіпаст. Напиши мені команду /help, щоб дізнатися більше.",
		Help: `Команди:
<code>/help</code> - повторно викликати це повідомлення.
<code>/addvoice &lt;назва&gt;</code> (в відповідь на аудіо/файл) - додати голосове повідомлення.
<code>/addpasta &lt;назва&gt;</code> (в відповідь на текст) - додати копіпасту.

Инлайн-команди:
<code>@%[1]s [v|voice|в|войс|гс] &lt;назва&gt;</code> - пошук голосового повідомлення.
<code>@%[1]s [p|pasta|п|паста] &lt;назва&gt;</code> - пошук копіпасти.`,

		NoName:     "Після команди необхідно написати назву.",
		NoReply:    "Цю команду можна використовувати тільки в відповідь на повідомлення.",
		VoiceAdded: "Голосове повідомлення додано.",
		PastaAdded: "Копіпаста додана.",

		CommandsList: "Список команд",
		NothingFound: "Нічого не знайдено",
	},

	"en": {
		Greeting: "Hi! I'm Labrys, a bot for sending voice messages and copypastas. Type /help, to learn more.",
		Help: `Commands:
<code>/help</code> - repeat this message.
<code>/addvoice &lt;name&gt;</code> (in reply to audio/file) - add voice message.
<code>/addpasta &lt;name&gt;</code> (in reply to text) - add copypasta.

Inline commands:
<code>@%[1]s [v|voice|в|войс|гс] &lt;name&gt;</code> - search voice message.
<code>@%[1]s [p|pasta|п|паста] &lt;name&gt;</code> - search copypasta.`,

		NoName:     "After the command, you need to type a name.",
		NoReply:    "This command can be used only in reply to message.",
		VoiceAdded: "Voice message added.",
		PastaAdded: "Copypasta added.",

		CommandsList: "Commands list",
		NothingFound: "Nothing found",
	},
}
