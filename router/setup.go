package router

import (
	"context"
	"crypto/subtle"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/app"
	"github.com/sparkymat/archmark/internal/handler"
	mw "github.com/sparkymat/archmark/middleware"
)

func Setup(e *echo.Echo, appService *app.Service) {
	e.Static("/css", "public/css")
	e.Static("/javascript", "public/javascript")
	e.Static("/b", appService.Config.DownloadPath())

	registerWebRoutes(e, appService)
	registerAPIRoutes(e, appService)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
	methodWhitelist := map[string]interface{}{
		"DELETE": struct{}{},
		"GET":    struct{}{},
		"PATCH":  struct{}{},
		"POST":   struct{}{},
		"PUT":    struct{}{},
	}

	fmt.Fprintf(os.Stdout, "\n  Registered routes  \n  =================  \n\n")

	for _, r := range e.Routes() {
		if _, whitelisted := methodWhitelist[r.Method]; whitelisted {
			if r.Path != "" && r.Path != "/*" {
				fmt.Fprintf(w, "%s\t%s\t\t%s\n", r.Method, r.Path, r.Name)
			}
		}
	}

	_ = w.Flush()
}

func registerWebRoutes(e *echo.Echo, appService *app.Service) {
	appGroup := e.Group("")

	appGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	appGroup.Use(middleware.Recover())
	appGroup.Use(mw.ConfigInjector(appService))
	appGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	appGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	appGroup.GET("/error", handler.ShowError)

	authAppGroup := appGroup.Group("")
	authAppGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("admin")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(appService.Config.AdminPassword())) == 1 {
			return true, nil
		}

		return false, nil
	}))
	authAppGroup.GET("/", handler.Home)
	authAppGroup.GET("/settings", handler.Settings)
	authAppGroup.POST("/settings", handler.UpdateSettings)
	authAppGroup.GET("/add", handler.BookmarksNew)
	authAppGroup.POST("/bookmarks", handler.BookmarksCreate)
	authAppGroup.POST("/bookmarks/:id/destroy", handler.BookmarksDestroy)
	authAppGroup.GET("/tokens", handler.APITokensIndex)
	authAppGroup.POST("/tokens/:id/destroy", handler.APITokensDestroy)
	authAppGroup.POST("/tokens", handler.APITokensCreate)
}

func registerAPIRoutes(e *echo.Echo, appService *app.Service) {
	apiAppGroup := e.Group("/api")
	apiAppGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	apiAppGroup.Use(middleware.Recover())
	apiAppGroup.Use(mw.ConfigInjector(appService))
	apiAppGroup.Use(middleware.KeyAuth(func(token string, c echo.Context) (bool, error) {
		_, err := appService.DB.LookupAPIToken(context.Background(), token)
		if err != nil {
			return false, fmt.Errorf("db lookup failed. err: %w", err)
		}

		return true, nil
	}))
	apiAppGroup.POST("/add", handler.APIBookmarksCreate)
}
