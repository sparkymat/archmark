package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/view"
)

func BookmarksNew(c echo.Context) error {
	pageHTML := view.BookmarksNew()
	htmlString := view.Layout("archmark | add", pageHTML)
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
