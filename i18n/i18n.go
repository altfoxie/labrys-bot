// Package i18n provides internationalization and localization support.
package i18n

// Locale represents localized messages.
type Locale struct {
	Greeting string
	Help     string

	NoName     string
	NoReply    string
	VoiceAdded string
	PastaAdded string

	NothingFound string
	CommandsList string
}

// Get returns locale for given language.
// If language is not supported, returns default locale.
func Get(lang string) *Locale {
	if _, ok := locales[lang]; !ok {
		lang = defaultLang
	}
	return locales[lang]
}
