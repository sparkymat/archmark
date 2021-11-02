package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/view"
)

func ShowError(c echo.Context) error {
	localizer := getLocalizer(c)
	if localizer == nil {
		localizer = localize.New()
	}

	lang := localize.English
	cfg := getConfig(c)
	if cfg != nil {
		lang = cfg.DefaultLanguage()
	}

	renderedError := localizer.Lookup(lang, localize.InternalServerError)

	return renderError(c, renderedError)
}

func renderError(c echo.Context, message string) error {
	localizer := getLocalizer(c)
	if localizer == nil {
		localizer = localize.New()
	}

	lang := localize.English
	cfg := getConfig(c)
	if cfg != nil {
		lang = cfg.DefaultLanguage()
	}

	pageHTML := view.ShowError(message)
	htmlString := view.Layout(localizer, lang, "archmark", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
