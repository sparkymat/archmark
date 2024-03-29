package handler

import (
	"context"

	"github.com/sparkymat/archmark/dbx"
)

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
	DisableRegistration() bool
	ProxyAuthUsernameHeader() string
	ProxyAuthNameHeader() string
}

type DatabaseService interface {
	FetchUserByUsername(ctx context.Context, email string) (dbx.User, error)
	FetchBookmarksList(ctx context.Context, arg dbx.FetchBookmarksListParams) ([]dbx.Bookmark, error)
	CountBookmarksList(ctx context.Context, userID int64) (int64, error)
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
}
