package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/view"
)

func BookmarksNew(c echo.Context) error {
	app := appServices(c)
	if app == nil {
		log.Print("error: app services not found")

		return ShowError(c)
	}

	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		return ShowError(c)
	}

	pageHTML := view.BookmarksNew(*app.Localizer, app.Settings.Language(), csrfToken)
	htmlString := view.Layout(*app.Localizer, app.Settings.Language(), "archmark | add", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
