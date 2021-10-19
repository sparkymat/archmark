package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/presenter"
	"github.com/sparkymat/archmark/view"
)

func ApiTokensIndex(c echo.Context) error {
	dbVal := c.Get(middleware.DBKey)
	if dbVal == nil {
		return c.String(http.StatusInternalServerError, "db conn not found")
	}
	db := dbVal.(database.API)

	tokens, err := db.ListApiTokens()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	presentedTokens := presenter.PresentApiTokens(tokens)
	pageHTML := view.ApiTokensIndex(presentedTokens)
	htmlString := view.Layout("archmark", pageHTML)
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
