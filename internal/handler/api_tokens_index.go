package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/database"
	mw "github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/presenter"
	"github.com/sparkymat/archmark/view"
)

func APITokensIndex(c echo.Context) error {
	csrfTokenVal := c.Get(middleware.DefaultCSRFConfig.ContextKey)
	if csrfTokenVal == nil {
		log.Print("error: csrf token not found")

		return ShowError(c)
	}

	csrfToken, ok := csrfTokenVal.(string)
	if !ok {
		log.Print("error: csrf token not found")

		return ShowError(c)
	}

	dbVal := c.Get(mw.DBKey)
	if dbVal == nil {
		log.Print("error: db conn not found")

		return ShowError(c)
	}

	db, ok := dbVal.(database.API)
	if !ok {
		log.Print("error: db conn not found")

		return ShowError(c)
	}

	tokens, err := db.ListAPITokens(c.Request().Context())
	if err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	presentedTokens := presenter.PresentAPITokens(tokens)
	pageHTML := view.ApiTokensIndex(csrfToken, presentedTokens)
	htmlString := view.Layout("archmark | tokens", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
