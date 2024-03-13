package internal

import (
	"context"

	"github.com/sparkymat/archmark/internal/dbx"
)

type UserService interface {
	FetchUserByEmail(ctx context.Context, email string) (dbx.User, error)
	CreateUser(ctx context.Context, name string, username string, encryptedPassword string) (dbx.User, error)
}
