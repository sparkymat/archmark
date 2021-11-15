package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/view"
)

func ShowError(c echo.Context) error {
	settings := getSettings(c)
	if settings == nil {
		log.Print("error: settings not found")

		return ShowError(c)
	}

	localizer := getLocalizer(c)
	if localizer == nil {
		log.Print("error: localizer not found")

		return ShowError(c)
	}

	renderedError := localizer.Lookup(settings.Language(), localize.InternalServerError)

	return renderError(c, renderedError)
}

func renderError(c echo.Context, message string) error {
	lang := localize.English

	settings := getSettings(c)
	if settings == nil {
		cfg := getConfig(c)
		if cfg != nil {
			lang = cfg.DefaultLanguage()
		}
	}

	localizer := getLocalizer(c)
	if localizer == nil {
		log.Print("error: localizer not found")
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", message))
	}

	pageHTML := view.ShowError(message)
	htmlString := view.Layout(localizer, lang, "archmark", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
