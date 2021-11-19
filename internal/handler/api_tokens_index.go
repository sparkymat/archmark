package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
	"github.com/sparkymat/archmark/presenter"
	"github.com/sparkymat/archmark/view"
)

func APITokensIndex(appService *app.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := getCSRFToken(c)
		if csrfToken == "" {
			log.Print("error: csrf token not found")

			return ShowError(appService)(c)
		}

		tokens, err := appService.DB.ListAPITokens(c.Request().Context())
		if err != nil {
			log.Printf("error: %v", err)

			return ShowError(appService)(c)
		}

		presentedTokens := presenter.PresentAPITokens(tokens)
		pageHTML := view.ApiTokensIndex(appService.Styler, appService.Localizer, appService.Settings.Language(), csrfToken, presentedTokens)
		htmlString := view.Layout(appService.Styler, appService.Localizer, appService.Settings.Language(), "archmark | tokens", pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
