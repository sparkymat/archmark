package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

type DatabaseProvider interface {
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
	FetchUserByEmail(ctx context.Context, email string) (dbx.User, error)
	FetchUserByID(ctx context.Context, userID uuid.UUID) (dbx.User, error)
}
