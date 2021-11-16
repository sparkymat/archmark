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

	app := appServices(c)
	if app == nil {
		log.Print("error: app services not found")

		return ShowError(c)
	}

	tokens, err := app.DB.ListAPITokens(c.Request().Context())
	if err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	presentedTokens := presenter.PresentAPITokens(tokens)
	pageHTML := view.ApiTokensIndex(*app.Localizer, app.Settings.Language(), csrfToken, presentedTokens)
	htmlString := view.Layout(*app.Localizer, app.Settings.Language(), "archmark | tokens", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
