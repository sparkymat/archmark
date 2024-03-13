package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/auth"
	"github.com/sparkymat/archmark/internal/handler"
)

func registerWebRoutes(app *echo.Group, cfg ConfigService, s internal.Services) {
	webApp := app.Group("")

	webApp.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	webApp.GET("/login", handler.Login(s))
	webApp.POST("/login", handler.DoLogin(s))

	if !cfg.DisableRegistration() {
		webApp.GET("/register", handler.Register(s))
		webApp.POST("/register", handler.DoRegister(s))
	}

	authenticatedWebApp := webApp.Group("")

	if cfg.ReverseProxyAuthentication() {
		authenticatedWebApp.Use(auth.ProxyAuthMiddleware(cfg, s.User))
	} else {
		authenticatedWebApp.Use(auth.Middleware(cfg, s.User))
	}

	authenticatedWebApp.GET("/", handler.Home(s))
}
