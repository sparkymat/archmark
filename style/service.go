package style

func New(theme Theme) *Service {
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
