package handler

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
)

func ApiTokensCreate(c echo.Context) error {
	dbVal := c.Get(middleware.DBKey)
	if dbVal == nil {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "db conn not found")
	}
	db := dbVal.(database.API)

	token, err := randomHex(32)
	if err != nil {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if _, err = db.CreateApiToken(token); err != nil {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, err.Error())
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, "/tokens")
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to read random data. err: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
