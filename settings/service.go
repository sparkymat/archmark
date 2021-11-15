package settings

import (
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/model"
)

type API interface {
	Language() localize.Language
}

func New(settings *model.Settings, cfg config.API) API {
	return &service{
		settings: settings,
		cfg:      cfg,
	}
}

type service struct {
	settings *model.Settings
	cfg      config.API
}

func (s *service) Language() localize.Language {
	if s.settings == nil {
		return s.cfg.DefaultLanguage()
	}

	return localize.LanguageFromString(s.settings.Language)
}
