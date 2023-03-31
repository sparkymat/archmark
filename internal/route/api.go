package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/auth"
)

func registerAPIRoutes(app *echo.Group, cfg ConfigService, db DatabaseService) {
	apiV2 := app.Group("api/v2")
	apiV2.Use(auth.APIMiddleware(cfg, db))

	apiV2.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token",
	}))
}
