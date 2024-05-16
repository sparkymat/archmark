package user

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal/dbx"
	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8
	defaultBcryptCost = 10
)

var (
	ErrPasswordTooShort      = errors.New("password too short")
	ErrUserAlreadyRegistered = errors.New("user already registered")
)

func (s *Service) Register(c echo.Context, name string, email string, password string, passwordConfirmation string) (dbx.User, error) {
	if password != passwordConfirmation {
		return dbx.User{}, ErrPasswordsDontMatch
	}

	if len(password) < minPasswordLength {
		return dbx.User{}, ErrPasswordTooShort
	}

	_, err := s.db.FetchUserByEmail(c.Request().Context(), email)
	if err == nil {
		return dbx.User{}, ErrUserAlreadyRegistered
	}

	encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultBcryptCost)
	if err != nil {
		return dbx.User{}, fmt.Errorf("failed to encrypt password: %w", err)
	}

	user, err := s.db.CreateUser(c.Request().Context(), dbx.CreateUserParams{
		Name:              name,
		Email:             email,
		EncryptedPassword: string(encryptedPasswordBytes),
	})
	if err != nil {
		return dbx.User{}, fmt.Errorf("failed to create new user: %w", err)
	}

	return user, nil
}
