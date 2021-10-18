package router

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/internal/handler"
	mw "github.com/sparkymat/archmark/middleware"
)

func Setup(e *echo.Echo, cfg config.API, db database.API) {
	e.Static("/css", "public/css")
	e.Static("/javascript", "public/javascript")

	app := e.Group("")
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(mw.ConfigInjector(cfg, db))
	app.POST("/add", handler.Create)

	authApp := app.Group("")
	authApp.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("admin")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(cfg.AdminPassword())) == 1 {
			return true, nil
		}

		return false, nil
	}))
	authApp.GET("/", handler.Home)
}
