package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/presenter"
	"github.com/sparkymat/archmark/view"
)

func APITokensIndex(c echo.Context) error {
	dbVal := c.Get(middleware.DBKey)
	if dbVal == nil {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "db conn not found")
	}

	db, ok := dbVal.(database.API)
	if !ok {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "db conn not found")
	}

	tokens, err := db.ListAPITokens(c.Request().Context())
	if err != nil {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, err.Error())
	}

	presentedTokens := presenter.PresentAPITokens(tokens)
	pageHTML := view.ApiTokensIndex(presentedTokens)
	htmlString := view.Layout("archmark | tokens", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
