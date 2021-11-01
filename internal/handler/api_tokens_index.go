package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/presenter"
	"github.com/sparkymat/archmark/view"
)

func APITokensIndex(c echo.Context) error {
	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		return ShowError(c)
	}

	db := getDB(c)
	if db == nil {
		log.Print("error: db conn not found")

		return ShowError(c)
	}

	localizer := getLocalizer(c)
	if localizer == nil {
		log.Print("error: localizer not found")

		return ShowError(c)
	}

	cfg := getConfig(c)
	if cfg == nil {
		log.Print("error: config not found")

		return ShowError(c)
	}

	tokens, err := db.ListAPITokens(c.Request().Context())
	if err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	presentedTokens := presenter.PresentAPITokens(tokens)
	pageHTML := view.ApiTokensIndex(csrfToken, presentedTokens)
	htmlString := view.Layout(localizer, cfg.DefaultLanguage(), "archmark | tokens", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
