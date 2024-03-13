package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/auth"
)

func registerAPIRoutes(app *echo.Group, cfg ConfigService, s internal.Services) {
	apiGroup := app.Group("api")

	if cfg.ReverseProxyAuthentication() {
		apiGroup.Use(auth.ProxyAuthMiddleware(cfg, s.User))
	} else {
		apiGroup.Use(auth.APIMiddleware(cfg, s.User))
	}

	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token",
	}))
}
