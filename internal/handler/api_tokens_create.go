package handler

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
)

const tokenLength = 32

func APITokensCreate(appService *app.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := randomHex(tokenLength)
		if err != nil {
			log.Printf("error: %v", err.Error())

			return ShowError(appService)(c)
		}

		if _, err = appService.DB.CreateAPIToken(c.Request().Context(), token); err != nil {
			log.Printf("error: %v", err.Error())

			return ShowError(appService)(c)
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/tokens")
	}
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to read random data. err: %w", err)
	}

	return hex.EncodeToString(bytes), nil
}
