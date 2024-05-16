package internal

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal/dbx"
)

type UserService interface {
	CreateUser(ctx context.Context, name string, username string, encryptedPassword string) (dbx.User, error)
	FetchUserByID(ctx context.Context, userID uuid.UUID) (dbx.User, error)
	FetchUserByEmail(ctx context.Context, email string) (dbx.User, error)
	Login(c echo.Context, email string, password string) error
	Register(c echo.Context, name string, email string, password string, passwordConfirmation string) (dbx.User, error)
}
