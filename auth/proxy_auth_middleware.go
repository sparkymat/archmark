package auth

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/dbx"
	"golang.org/x/crypto/bcrypt"
)

const defaultBcryptCost = 10

func ProxyAuthMiddleware(cfg ConfigService, db DatabaseService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			name := c.Request().Header.Get(cfg.ProxyAuthNameHeader())

			username := c.Request().Header.Get(cfg.ProxyAuthUsernameHeader())
			if username == "" {
				return c.String(http.StatusUnauthorized, "user header missing")
			}

			user, err := db.FetchUserByUsername(c.Request().Context(), username)
			if err == nil {
				c.Set(UserKey, user)

				return next(c)
			}

			password := strings.ReplaceAll(uuid.New().String(), "-", "")

			encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultBcryptCost)
			if err != nil {
				return c.String(http.StatusUnauthorized, "failed to generate password")
			}

			user, err = db.CreateUser(c.Request().Context(), dbx.CreateUserParams{
				Name:              name,
				Username:          username,
				EncryptedPassword: string(encryptedPasswordBytes),
			})
			if err != nil {
				return c.String(http.StatusUnauthorized, "failed to add new user")
			}

			c.Set(UserKey, user)

			return next(c)
		}
	}
}
