package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/model"
)

type UpdateSettingsInput struct {
	Language string `form:"language"`
}

func UpdateSettings(c echo.Context) error {
	var input UpdateSettingsInput

	if err := c.Bind(&input); err != nil {
		log.Printf("error: %v", err)

		return renderError(c, "Unable to update settings. Please try again later.")
	}

	app := appServices(c)
	if app == nil {
		log.Print("error: app services not found")

		return ShowError(c)
	}

	settingsModel, err := app.DB.LoadSettings(c.Request().Context(), model.DefaultSettings(*app.Config))
	if err != nil {
		log.Print("error: failed to load settings")

		return ShowError(c)
	}

	settingsModel.Language = input.Language

	err = app.DB.UpdateSettings(c.Request().Context(), settingsModel)
	if err != nil {
		log.Print("error: failed to update settings")

		return ShowError(c)
	}

	err = app.RefreshSettings(c.Request().Context())
	if err != nil {
		log.Print("error: failed to refresh settings")

		return ShowError(c)
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, "/settings")
}
