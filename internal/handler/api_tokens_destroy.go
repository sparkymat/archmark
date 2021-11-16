package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	base10        = 10
	sixtyFourBits = 64
)

func APITokensDestroy(c echo.Context) error {
	app := appServices(c)
	if app == nil {
		log.Print("error: app services not found")

		return ShowError(c)
	}

	tokenIDString := c.Param("id")

	tokenID, err := strconv.ParseUint(tokenIDString, base10, sixtyFourBits)
	if err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	if err = app.DB.DeleteAPIToken(c.Request().Context(), tokenID); err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, "/tokens")
}
