package handler

import (
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
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "db conn not found")
	}

	db, ok := dbVal.(database.API)
	if !ok {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "db conn not found")
	}

	tokenIDString := c.Param("id")

	tokenID, err := strconv.ParseUint(tokenIDString, base10, sixtyFourBits)
	if err != nil {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err = db.DeleteAPIToken(uint(tokenID)); err != nil {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, err.Error())
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, "/tokens")
}
