package route

import (
	"context"

	"github.com/sparkymat/archmark/dbx"
)

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
	DatabaseURL() string
}

type DatabaseService interface {
	FetchUserByUsername(ctx context.Context, email string) (dbx.User, error)
	FetchBookmarksList(ctx context.Context, arg dbx.FetchBookmarksListParams) ([]dbx.Bookmark, error)
	CountBookmarksList(ctx context.Context, userID int64) (int64, error)
	CreateBookmark(ctx context.Context, arg dbx.CreateBookmarkParams) (dbx.Bookmark, error)
}
