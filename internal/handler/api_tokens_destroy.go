package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
)

const (
	base10        = 10
	sixtyFourBits = 64
)

func APITokensDestroy(appService *app.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenIDString := c.Param("id")

		tokenID, err := strconv.ParseUint(tokenIDString, base10, sixtyFourBits)
		if err != nil {
			log.Printf("error: %v", err)

			return ShowError(appService)(c)
		}

		if err = appService.DB.DeleteAPIToken(c.Request().Context(), tokenID); err != nil {
			log.Printf("error: %v", err)

			return ShowError(appService)(c)
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/tokens")
	}
}
