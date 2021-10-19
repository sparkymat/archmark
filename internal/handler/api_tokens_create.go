package handler

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
)

func ApiTokensCreate(c echo.Context) error {
	dbVal := c.Get(middleware.DBKey)
	if dbVal == nil {
		return c.String(http.StatusInternalServerError, "db conn not found")
	}
	db := dbVal.(database.API)

	token, err := randomHex(32)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	_, err = db.CreateApiToken(token)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/tokens")

}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
