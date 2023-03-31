package handler

import (
	"context"

	"github.com/sparkymat/archmark/dbx"
)

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
}

type DatabaseService interface {
	FetchUserByUsername(ctx context.Context, email string) (dbx.User, error)
}
