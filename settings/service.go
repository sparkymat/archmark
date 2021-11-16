package settings

import (
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/model"
)

func New(settings *model.Settings, cfg *config.Service) Service {
	return Service{
		settings: settings,
		cfg:      cfg,
	}
}

type Service struct {
	settings *model.Settings
	cfg      *config.Service
}

func (s *Service) Language() localize.Language {
	if s.settings == nil {
		if s.cfg == nil {
			return localize.English
		}

		return s.cfg.DefaultLanguage()
	}

	return localize.LanguageFromString(s.settings.Language)
}
