package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
	"github.com/sparkymat/archmark/view"
)

func BookmarksNew(appService *app.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := getCSRFToken(c)
		if csrfToken == "" {
			log.Print("error: csrf token not found")

			return ShowError(appService)(c)
		}

		pageHTML := view.BookmarksNew(appService.Styler.Theme(), appService.Localizer, appService.Settings.Language(), csrfToken)
		htmlString := view.Layout(appService.Styler.Theme(), appService.Localizer, appService.Settings.Language(), "archmark | add", pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
