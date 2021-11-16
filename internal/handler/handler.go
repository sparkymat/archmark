package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/localize"
	mw "github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/settings"
)

func getSettings(c echo.Context) settings.API {
	settingsVal := c.Get(mw.SettingsKey)
	if settingsVal == nil {
		return nil
	}

	settingsService, ok := settingsVal.(settings.API)
	if !ok {
		return nil
	}

	return settingsService
}

func getCSRFToken(c echo.Context) string {
	csrfTokenVal := c.Get(middleware.DefaultCSRFConfig.ContextKey)
	if csrfTokenVal == nil {
		return ""
	}

	csrfToken, ok := csrfTokenVal.(string)
	if !ok {
		return ""
	}

	return csrfToken
}

func getDB(c echo.Context) database.API {
	dbVal := c.Get(mw.DBKey)
	if dbVal == nil {
		return nil
	}

	db, ok := dbVal.(database.API)
	if !ok {
		return nil
	}

	return db
}

func getConfig(c echo.Context) config.API {
	configVal := c.Get(mw.ConfigKey)
	if configVal == nil {
		return nil
	}

	cfg, ok := configVal.(config.API)
	if !ok {
		return nil
	}

	return cfg
}

func getLocalizer(c echo.Context) *localize.Service {
	localizeVal := c.Get(mw.LocalizeKey)
	if localizeVal == nil {
		return nil
	}

	localizer, ok := localizeVal.(*localize.Service)
	if !ok {
		return nil
	}

	return localizer
}
