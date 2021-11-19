package app

import (
	"context"
	"fmt"

	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/settings"
	"github.com/sparkymat/archmark/style"
)

type Service struct {
	Config    *config.Service
	DB        *database.Service
	Localizer *localize.Service
	Settings  *settings.Service
	Styler    *style.Service
}

func New(configService *config.Service, db *database.Service, localizer *localize.Service, settingsService *settings.Service, stylerService *style.Service) *Service {
	return &Service{
		Config:    configService,
		DB:        db,
		Localizer: localizer,
		Settings:  settingsService,
		Styler:    stylerService,
	}
}

func (s *Service) RefreshSettings(ctx context.Context) error {
	settingsModel, err := s.DB.LoadSettings(ctx, model.DefaultSettings(s.Config))
	if err != nil {
		return fmt.Errorf("failed to load settings. err: %w", err)
	}

	s.Settings = settings.New(settingsModel, s.Config)

	return nil
}
