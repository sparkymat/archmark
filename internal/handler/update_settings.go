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

		return renderError(c, "Unable to add bookmark. Please try again later.")
	}

	return c.Redirect(http.StatusSeeOther, "/settings")
}
