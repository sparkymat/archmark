package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
)

type ContextKey string

const (
	AppServicesKey = "app"
)

func ConfigInjector(appService *app.Service) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(AppServicesKey, appService)

			return next(c)
		}
	}
}
