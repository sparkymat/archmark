package handler

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const tokenLength = 32

func APITokensCreate(c echo.Context) error {
	db := getDB(c)
	if db == nil {
		log.Print("error: db conn not found")

		return ShowError(c)
	}

	token, err := randomHex(tokenLength)
	if err != nil {
		log.Printf("error: %v", err.Error())

		return ShowError(c)
	}

	if _, err = db.CreateAPIToken(c.Request().Context(), token); err != nil {
		log.Printf("error: %v", err.Error())

		return ShowError(c)
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
