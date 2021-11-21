package style

const (
	ThemeLight = "light"
	ThemeDark  = "dark"
)

func New(themeName string) *Service {
	var theme Theme

	themeFn, themeFound := map[string]func() Theme{
		ThemeLight: lightTheme,
		ThemeDark:  darkTheme,
	}[themeName]
	if themeFound {
		theme = themeFn()
	} else {
		theme = lightTheme()
	}

	return &Service{
		theme: theme,
	}
}

type Service struct {
	theme Theme
}

func (s *Service) Theme() Theme {
	return s.theme
}

func (s *Service) SetTheme(theme Theme) {
	s.theme = theme
}
