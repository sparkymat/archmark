package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/auth"
	"github.com/sparkymat/archmark/internal/handler"
)

func registerWebRoutes(app *echo.Group, cfg ConfigService, db DatabaseService) {
	webApp := app.Group("")

	webApp.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	webApp.GET("/login", handler.Login(cfg, db))
	webApp.POST("/login", handler.DoLogin(cfg, db))

	authenticatedWebApp := webApp.Group("")
	authenticatedWebApp.Use(auth.Middleware(cfg, db))

	authenticatedWebApp.GET("/", handler.Home())
}