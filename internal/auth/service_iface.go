package auth

import (
	"context"

	"github.com/sparkymat/archmark/internal/dbx"
)

type ConfigService interface {
	JWTSecret() string
	ProxyAuthUsernameHeader() string
	ProxyAuthNameHeader() string
}

type UserService interface {
	FetchUserByEmail(ctx context.Context, email string) (dbx.User, error)
	CreateUser(ctx context.Context, name string, email string, encryptedPassword string) (dbx.User, error)
}
