package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/service/user"
	"github.com/sparkymat/archmark/internal/view"
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

		if _, err := s.User.Register(c, name, email, password, passwordConfirmation); err != nil {
			switch err {
			case user.ErrPasswordsDontMatch:
				return renderRegistrationPage(c, name, email, "Passwords don't match")
			case user.ErrPasswordTooShort:
				return renderRegistrationPage(c, name, email, "Password too short")
			case user.ErrUserAlreadyRegistered:
				return renderRegistrationPage(c, name, email, "User already registered")
			default:
				return renderRegistrationPage(c, name, email, "Registration failed")
			}
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
