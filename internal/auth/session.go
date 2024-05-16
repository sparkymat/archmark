package auth

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var ErrSessionError = errors.New("session error")

func LoadUserIDFromSession(jwtSecret string, c echo.Context) (uuid.UUID, error) {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return uuid.Nil, err //nolint:wrapcheck
	}

	tokenInterface, valueExists := sess.Values[tokenKey]
	if !valueExists {
		return uuid.Nil, ErrSessionError
	}

	token, isString := tokenInterface.(string)
	if !isString {
		return uuid.Nil, ErrSessionError
	}

	userIDString, err := ParseJWTString(jwtSecret, token)
	if err != nil {
		return uuid.Nil, err
	}

	if userIDString == nil {
		return uuid.Nil, ErrSessionError
	}

	userID, err := uuid.Parse(*userIDString)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

func SaveUserIDToSession(jwtSecret string, c echo.Context, userID uuid.UUID) error {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return fmt.Errorf("failed to initialize session. err: %w", err)
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 8, //nolint:gomnd
		HttpOnly: true,
	}

	token, err := GenerateJWTString(jwtSecret, userID)
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
