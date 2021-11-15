package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/settings"
)

func SettingsInjector(cfg config.API, db database.API) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			settingsModel, err := db.LoadSettings(c.Request().Context(), model.DefaultSettings(cfg))
			if err != nil {
				log.Printf("failed to load settings. err: %v\n", err)
				return err
			}

			settingsService := settings.New(settingsModel, cfg)

			c.Set(SettingsKey, settingsService)

			return next(c)
		}
	}
}
