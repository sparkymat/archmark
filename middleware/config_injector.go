package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
)

type ContextKey string

const (
	ConfigKey = "config"
	DBKey     = "db"
)

func ConfigInjector(cfg config.API, db database.API) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(ConfigKey, cfg)
			c.Set(DBKey, db)
			return next(c)
		}
	}
}
