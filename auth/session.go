package auth

import (
	"errors"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/jwt"
)

var ErrSessionError = errors.New("session error")

func LoadUsernameFromSession(cfg ConfigService, c echo.Context) (string, error) {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return "", err //nolint:wrapcheck
	}

	tokenInterface, valueExists := sess.Values[tokenKey]
	if !valueExists {
		return "", ErrSessionError
	}

	token, isString := tokenInterface.(string)
	if !isString {
		return "", ErrSessionError
	}

	email, err := jwt.ValidateTokenString(cfg.JWTSecret(), token)
	if err != nil {
		return "", err //nolint:wrapcheck
	}

	if email == nil {
		return "", ErrSessionError
	}

	return *email, nil
}

func SaveUsernameToSession(cfg ConfigService, c echo.Context, username string) error {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return fmt.Errorf("failed to initialize session. err: %w", err)
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 8, //nolint:gomnd
		HttpOnly: true,
	}

	token, err := jwt.GenerateToken(cfg.JWTSecret(), username)
	if err != nil {
		return fmt.Errorf("failed to generate token. err: %w", err)
	}

	sess.Values[tokenKey] = token

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return fmt.Errorf("failed to save session. err: %w", err)
	}

	return nil
}
