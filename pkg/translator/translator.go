package translator

import (
	"strings"
)

type (
	Translator interface {
		Translate(key string, lang ...Language) string
	}

	Language string
)

const (
	EN Language = "en"
	FA Language = "fa"
)

func GetLanguage(lang string) Language {
	switch strings.ToLower(lang) {
	case "en", "en-us":
		return EN
	case "fa", "fa-ir":
		return FA
	default:
		return EN
	}
}
