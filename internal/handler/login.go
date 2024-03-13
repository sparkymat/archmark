package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/auth"
	"github.com/sparkymat/archmark/internal/view"
	"golang.org/x/crypto/bcrypt"
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

		user, err := s.User.FetchUserByEmail(c.Request().Context(), email)
		if err != nil {
			log.Printf("failed to load user. err: %v", err)

			return renderLoginPage(s, c, email, "Authentication failed")
		}

		if bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) != nil {
			log.Printf("password does not match")

			return renderLoginPage(s, c, email, "Authentication failed")
		}

		err = auth.SaveEmailToSession(s, c, user.Email)
		if err != nil {
			log.Printf("failed to save email to session. err: %v", err)

			return renderLoginPage(s, c, email, "Authentication failed")
		}

		//nolint:wrapcheck
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

	pageHTML := view.Login(cfg.DisableRegistration(), csrfToken, email, errorMessage)
	htmlString := view.Layout("oxbook | login", csrfToken, pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
