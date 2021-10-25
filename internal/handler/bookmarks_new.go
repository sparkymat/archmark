package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/view"
)

func BookmarksNew(c echo.Context) error {
	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		return ShowError(c)
	}

	pageHTML := view.BookmarksNew(csrfToken)
	htmlString := view.Layout("archmark | add", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
