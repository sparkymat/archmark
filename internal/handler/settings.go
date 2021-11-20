package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/presenter"
	"github.com/sparkymat/archmark/view"
)

func Settings(appService *app.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := getCSRFToken(c)
		if csrfToken == "" {
			log.Print("error: csrf token not found")

			return ShowError(appService)(c)
		}

		presentedLanguages := presenter.SupportedLanguages(localize.SupportedLanguages)

		settingsModel, err := appService.DB.LoadSettings(c.Request().Context(), model.DefaultSettings(appService.Config))
		if err != nil {
			log.Print("error: settings not found")

			return ShowError(appService)(c)
		}

		pageHTML := view.Settings(appService.Localizer, appService.Settings.Language(), csrfToken, presentedLanguages, settingsModel.Language)
		htmlString := view.Layout(appService.Styler.Theme(), appService.Localizer, appService.Settings.Language(), "archmark", pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
