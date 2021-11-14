package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/model"
)

func SettingsInjector(cfg config.API, db database.API) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			settings, err := db.LoadSettings(c.Request().Context())
			if err != nil {
				defaultSettings := model.DefaultSettings(cfg)
				settings = &defaultSettings
			}

			c.Set(SettingsKey, settings)

			return next(c)
		}
	}
}
