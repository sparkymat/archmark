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
	app := appServices(c)
	if app == nil {
		log.Print("error: app services not found")

		return ShowError(c)
	}

	renderedError := app.Localizer.Lookup(app.Settings.Language(), localize.InternalServerError)

	return renderError(c, renderedError)
}

func renderError(c echo.Context, message string) error {
	lang := localize.English

	app := appServices(c)
	if app == nil {
		log.Print("error: app services not found")

		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", message))
	}

	pageHTML := view.ShowError(message)
	htmlString := view.Layout(*app.Localizer, lang, "archmark", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
