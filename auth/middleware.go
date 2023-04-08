package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/dbx"
)

const (
	UserKey = "user"
)

const (
	sessionName = "archmark_session"
	tokenKey    = "auth_token"
)

var ErrTokenMissing = errors.New("token missing")

type ConfigService interface {
	JWTSecret() string
	ProxyAuthUsernameHeader() string
	ProxyAuthNameHeader() string
}

type DatabaseService interface {
	FetchUserByUsername(ctx context.Context, username string) (dbx.User, error)
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
}

const ClientNameKey = "client_name"

func Middleware(cfg ConfigService, db DatabaseService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := sessionAuth(cfg, db, c)
			if err == nil {
				return next(c)
			}

			return c.Redirect(http.StatusSeeOther, "/login") //nolint:wrapcheck
		}
	}
}

func APIMiddleware(cfg ConfigService, db DatabaseService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := sessionAuth(cfg, db, c)
			if err == nil {
				return next(c)
			}

			return c.Redirect(http.StatusSeeOther, "/login") //nolint:wrapcheck
		}
	}
}

func sessionAuth(cfg ConfigService, db DatabaseService, c echo.Context) error {
	username, err := LoadUsernameFromSession(cfg, c)
	if err != nil {
		return err
	}

	user, err := db.FetchUserByUsername(c.Request().Context(), username)
	if err != nil {
		return err //nolint:wrapcheck
	}

	c.Set(UserKey, user)

	return nil
}
