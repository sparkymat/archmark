package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
	"github.com/sparkymat/archmark/model"
)

type UpdateSettingsInput struct {
	Language string `form:"language"`
}

func UpdateSettings(appService *app.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateSettingsInput

		if err := c.Bind(&input); err != nil {
			log.Printf("error: %v", err)

			return renderError(c, appService, "Unable to update settings. Please try again later.")
		}

		settingsModel, err := appService.DB.LoadSettings(c.Request().Context(), model.DefaultSettings(appService.Config))
		if err != nil {
			log.Print("error: failed to load settings")

			return ShowError(appService)(c)
		}

		settingsModel.Language = input.Language

		err = appService.DB.UpdateSettings(c.Request().Context(), settingsModel)
		if err != nil {
			log.Print("error: failed to update settings")

			return ShowError(appService)(c)
		}

		err = appService.RefreshSettings(c.Request().Context())
		if err != nil {
			log.Print("error: failed to refresh settings")

			return ShowError(appService)(c)
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/settings")
	}
}
