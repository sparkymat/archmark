package router

import (
	"context"
	"crypto/subtle"
	"fmt"
	"os"
	"text/tabwriter"

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
	e.Static("/b", cfg.DownloadPath())

	app := e.Group("")

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.Recover())
	app.Use(mw.ConfigInjector(cfg, db))
	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

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
	authApp.GET("/add", handler.BookmarksNew)
	authApp.POST("/bookmarks", handler.BookmarksCreate)
	authApp.GET("/tokens", handler.APITokensIndex)
	authApp.POST("/tokens/:id/destroy", handler.APITokensDestroy)
	authApp.POST("/tokens", handler.APITokensCreate)

	apiApp := e.Group("/api")
	apiApp.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	apiApp.Use(middleware.Recover())
	apiApp.Use(mw.ConfigInjector(cfg, db))
	apiApp.Use(middleware.KeyAuth(func(token string, c echo.Context) (bool, error) {
		_, err := db.LookupAPIToken(context.Background(), token)
		if err != nil {
			return false, fmt.Errorf("db lookup failed. err: %w", err)
		}

		return true, nil
	}))
	apiApp.POST("/add", handler.APIBookmarksCreate)

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

	w.Flush()
}
