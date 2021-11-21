package presenter

import (
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/style"
)

type Theme struct {
	Label string
	Value string
}

func SupportedThemes(localizer *localize.Service, lang localize.Language) []Theme {
	supportedThemeStrngIdentifiers := map[string]localize.StringIdentifier{
		style.ThemeLight: localize.ThemeLight,
		style.ThemeDark:  localize.ThemeDark,
	}

	themes := []Theme{}
	for themeValue, themeStringIdentifier := range supportedThemeStrngIdentifiers {
		themes = append(themes, Theme{
			Label: localizer.Lookup(lang, themeStringIdentifier),
			Value: themeValue,
		})
	}

	return themes
}
