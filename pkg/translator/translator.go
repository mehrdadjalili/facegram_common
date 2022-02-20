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
	AR Language = "ar"
	RU Language = "ru"
	TR Language = "tr"
	ZH Language = "zh"
	JA Language = "ja"
	KO Language = "ko"
	DE Language = "de"
	ES Language = "es"
	FR Language = "fr"
	HI Language = "hi"
)

func GetLanguage(lang string) Language {
	switch strings.ToLower(lang) {
	case "en", "en-us", "en-US":
		return EN
	case "fa", "fa-ir", "fa-IR":
		return FA
	case "ar", "ar-ab", "ar-AB":
		return AR
	case "ru", "ru-ru", "ru-RU":
		return RU
	case "tr", "tr-tr", "tr-TR":
		return TR
	case "zh", "zh-hk", "zh-HK":
		return ZH
	case "ja", "ja-ja", "ja-JA":
		return JA
	case "ko", "ko-kr", "ko-KR":
		return KO
	case "de", "de-de", "de-DE":
		return DE
	case "es", "es-es", "es-ES":
		return ES
	case "fr", "fr-fr", "fr-FR":
		return FR
	case "hi", "hi-in", "hi-HI":
		return HI
	default:
		return EN
	}
}
