package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/view"
	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8
	defaultBcryptCost = 10
)

func Register(_ internal.Services) echo.HandlerFunc {
	return func(c echo.Context) error {
		return renderRegistrationPage(c, "", "", "")
	}
}

func DoRegister(s internal.Services) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		passwordConfirmation := c.FormValue("password_confirmation")

		if password != passwordConfirmation {
			log.Printf("passwords don't match")

			return renderRegistrationPage(c, name, email, "Passwords don't match")
		}

		if len(password) < minPasswordLength {
			log.Printf("password too short")

			return renderRegistrationPage(c, name, email, "Password too short")
		}

		_, err := s.User.FetchUserByEmail(c.Request().Context(), email)
		if err == nil {
			return renderRegistrationPage(c, name, email, "User already registered")
		}

		encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultBcryptCost)
		if err != nil {
			log.Printf("password does not match")

			return renderRegistrationPage(c, name, email, "Registration failed")
		}

		_, err = s.User.CreateUser(c.Request().Context(), name, email, string(encryptedPasswordBytes))
		if err != nil {
			log.Printf("failed to create new user. err: %v", err)

			return renderRegistrationPage(c, name, email, "Registration failed")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/login")
	}
}

func renderRegistrationPage(c echo.Context, name string, username string, errorMessage string) error {
	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "server error")
	}

	pageHTML := view.Register(csrfToken, name, username, errorMessage)
	htmlString := view.Layout("archmark | register", csrfToken, pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
