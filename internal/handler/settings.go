package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/presenter"
	"github.com/sparkymat/archmark/view"
)

func Settings(c echo.Context) error {
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

	presentedLanguages := presenter.SupportedLanguages(localize.SupportedLanguages)

	settingsModel, err := app.DB.LoadSettings(c.Request().Context(), model.DefaultSettings(app.Config))
	if err != nil {
		log.Print("error: settings not found")

		return ShowError(c)
	}

	pageHTML := view.Settings(*app.Localizer, app.Settings.Language(), csrfToken, presentedLanguages, settingsModel.Language)
	htmlString := view.Layout(*app.Localizer, app.Settings.Language(), "archmark", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
