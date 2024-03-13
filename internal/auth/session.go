package auth

import (
	"errors"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var ErrSessionError = errors.New("session error")

func LoadEmailFromSession(cfg ConfigService, c echo.Context) (string, error) {
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

	email, err := ValidateJWTString(cfg.JWTSecret(), token)
	if err != nil {
		return "", err
	}

	if email == nil {
		return "", ErrSessionError
	}

	return *email, nil
}

func SaveEmailToSession(cfg ConfigService, c echo.Context, email string) error {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return fmt.Errorf("failed to initialize session: %w", err)
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 8, //nolint:gomnd
		HttpOnly: true,
	}

	token, err := GenerateJWTString(cfg.JWTSecret(), email)
	if err != nil {
		return fmt.Errorf("failed to generate token: %w", err)
	}

	sess.Values[tokenKey] = token

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}
