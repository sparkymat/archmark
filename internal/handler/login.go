package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/service/features"
	"github.com/sparkymat/archmark/internal/view"
)

func Login(s internal.Services) echo.HandlerFunc {
	return func(c echo.Context) error {
		return renderLoginPage(s, c, "", "")
	}
}

func DoLogin(s internal.Services) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		err := s.User.Login(c, email, password)
		if err != nil {
			log.Printf("failed to log user in: %s", err.Error())

			return renderLoginPage(s, c, email, "Authentication failed")
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}
}

func renderLoginPage(s internal.Services, c echo.Context, email string, errorMessage string) error {
	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "server error")
	}

	pageHTML := view.Login(s.Features.Enabled(features.FeatureFlagRegistrationDisabled), csrfToken, email, errorMessage)
	htmlString := view.Layout("archmark | login", csrfToken, pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
