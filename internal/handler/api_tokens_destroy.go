package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
)

const (
	base10        = 10
	sixtyFourBits = 64
)

func APITokensDestroy(c echo.Context) error {
	dbVal := c.Get(middleware.DBKey)
	if dbVal == nil {
		fmt.Print("error: db conn not found")
		//nolint:wrapcheck
		return ShowError(c)
	}

	db, ok := dbVal.(database.API)
	if !ok {
		fmt.Print("error: db conn not found")
		//nolint:wrapcheck
		return ShowError(c)
	}

	tokenIDString := c.Param("id")

	tokenID, err := strconv.ParseUint(tokenIDString, base10, sixtyFourBits)
	if err != nil {
		log.Printf("error: %v", err)
		//nolint:wrapcheck
		return ShowError(c)
	}

	if err = db.DeleteAPIToken(c.Request().Context(), tokenID); err != nil {
		log.Printf("error: %v", err)
		//nolint:wrapcheck
		return ShowError(c)
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, "/tokens")
}
