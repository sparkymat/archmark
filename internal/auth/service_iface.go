package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

type ConfigService interface {
	JWTSecret() string
	ProxyAuthEmailHeader() string
	ProxyAuthNameHeader() string
}

type UserService interface {
	CreateUser(ctx context.Context, name string, email string, encryptedPassword string) (dbx.User, error)
	FetchUserByEmail(ctx context.Context, email string) (dbx.User, error)
	FetchUserByID(ctx context.Context, userID uuid.UUID) (dbx.User, error)
}
