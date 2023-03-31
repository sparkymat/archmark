package route

import (
	"context"

	"github.com/sparkymat/archmark/dbx"
)

type DatabaseService interface {
	FetchUserByUsername(ctx context.Context, email string) (dbx.User, error)
}

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
	DatabaseURL() string
}
