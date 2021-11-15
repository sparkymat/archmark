package presenter

import "github.com/sparkymat/archmark/localize"

type Language struct {
	Label string
	Value string
}

func SupportedLanguages(languageCodes []localize.Language) []Language {
	languages := []Language{}

	languageLabelMap := map[localize.Language]string{
		localize.English:   "English",
		localize.German:    "Deutsch",
		localize.Malayalam: "മലയാളം",
	}

	for _, languageCode := range languageCodes {
		if label, ok := languageLabelMap[languageCode]; ok {
			languages = append(languages, Language{
				Label: label,
				Value: string(languageCode),
			})
		}
	}

	return languages
}
