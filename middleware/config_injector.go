package middleware

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/settings"
)

type ContextKey string

const (
	AppServicesKey = "app_services"
)

type AppServices struct {
	Config    *config.Service
	DB        *database.Service
	Localizer *localize.Service
	Settings  settings.API
}

func ConfigInjector(appServices AppServices) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(AppServicesKey, &appServices)

			return next(c)
		}
	}
}

func (app *AppServices) RefreshSettings(ctx context.Context) error {
	settingsModel, err := app.DB.LoadSettings(ctx, model.DefaultSettings(*app.Config))
	if err != nil {
		return fmt.Errorf("failed to load settings. err: %w", err)
	}

	app.Settings = settings.New(settingsModel, *app.Config)

	return nil
}
