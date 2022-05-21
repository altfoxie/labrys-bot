// Package i18n provides internationalization and localization support.
package i18n

// Locale represents localized messages.
type Locale struct {
	Greeting           string
	UnknownMessage     string
	CommandsList       string
	InlineCommandsHelp string
	NothingFound       string
}

// Get returns locale for given language.
// If language is not supported, returns default locale.
func Get(lang string) *Locale {
	if _, ok := locales[lang]; !ok {
		lang = defaultLang
	}
	return locales[lang]
}
