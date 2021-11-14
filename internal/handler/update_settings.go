package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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

	db := getDB(c)
	if db == nil {
		log.Print("error: db not found")

		return ShowError(c)
	}

	settings := getSettings(c)
	if settings == nil {
		log.Print("error: settings not found")

		return ShowError(c)
	}

	settings.Language = input.Language

	err := db.UpdateSettings(c.Request().Context(), settings)
	if err != nil {
		log.Print("error: failed to update settings")

		return ShowError(c)
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, "/settings")
}
