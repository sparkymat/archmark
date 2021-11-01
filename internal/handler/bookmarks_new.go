package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/view"
)

func BookmarksNew(c echo.Context) error {
	localizer := getLocalizer(c)
	if localizer == nil {
		log.Print("error: localizer not found")

		return ShowError(c)
	}

	cfg := getConfig(c)
	if cfg == nil {
		log.Print("error: config not found")

		return ShowError(c)
	}

	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		return ShowError(c)
	}

	pageHTML := view.BookmarksNew(csrfToken)
	htmlString := view.Layout(localizer, cfg.DefaultLanguage(), "archmark | add", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
