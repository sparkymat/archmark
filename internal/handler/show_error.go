package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/view"
)

func ShowError(appService *app.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		renderedError := appService.Localizer.Lookup(appService.Settings.Language(), localize.InternalServerError)

		return renderError(c, appService, renderedError)
	}
}

func renderError(c echo.Context, appService *app.Service, message string) error {
	lang := localize.English
	pageHTML := view.ShowError(message)
	htmlString := view.Layout(appService.Styler, appService.Localizer, lang, "archmark", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
