package user

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

var ErrPasswordsDontMatch = errors.New("password does not match")

func (s *Service) Login(c echo.Context, email string, password string) error {
	user, err := s.db.FetchUserByEmail(c.Request().Context(), email)
	if err != nil {
		return fmt.Errorf("failed to load user: %w", err)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) != nil {
		return ErrPasswordsDontMatch
	}

	err = auth.SaveUserIDToSession(s.config.JWTSecret(), c, user.ID)
	if err != nil {
		return fmt.Errorf("failed to save email to session: %w", err)
	}

	return nil
}
