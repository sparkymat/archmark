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

	db := getDB(c)
	if db == nil {
		log.Print("error: db not found")

		return ShowError(c)
	}

	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		return ShowError(c)
	}

	settings := getSettings(c)
	if settings == nil {
		log.Print("error: settings not found")

		return ShowError(c)
	}

	presentedLanguages := presenter.SupportedLanguages(localize.SupportedLanguages)
	settingsModel, err := db.LoadSettings(c.Request().Context(), model.DefaultSettings(cfg))
	if err != nil {
		log.Print("error: settings not found")

		return ShowError(c)
	}

	pageHTML := view.Settings(localizer, settings.Language(), csrfToken, presentedLanguages, settingsModel.Language)
	htmlString := view.Layout(localizer, settings.Language(), "archmark", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
