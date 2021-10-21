package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
)

func ApiTokensDestroy(c echo.Context) error {
	dbVal := c.Get(middleware.DBKey)
	if dbVal == nil {
		return c.String(http.StatusInternalServerError, "db conn not found")
	}
	db := dbVal.(database.API)

	tokenIDString := c.Param("id")
	tokenID, err := strconv.ParseUint(tokenIDString, 10, 32)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err = db.DeleteApiToken(uint(tokenID)); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, "/tokens")
}
