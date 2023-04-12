package api

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sparkymat/archmark/dbx"
)

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
}

type DatabaseService interface {
	FetchUserByUsername(ctx context.Context, email string) (dbx.User, error)
	FetchBookmarksList(ctx context.Context, arg dbx.FetchBookmarksListParams) ([]dbx.Bookmark, error)
	CountBookmarksList(ctx context.Context, userID int64) (int64, error)
	CreateBookmark(ctx context.Context, arg dbx.CreateBookmarkParams) (dbx.Bookmark, error)
	SearchBookmarks(ctx context.Context, arg dbx.SearchBookmarksParams) ([]dbx.Bookmark, error)
	CountBookmarksSearchResults(ctx context.Context, arg dbx.CountBookmarksSearchResultsParams) (int64, error)
	FetchCategories(ctx context.Context, userID int64) ([]pgtype.Text, error)
	UpdateBookmarkCategory(ctx context.Context, arg dbx.UpdateBookmarkCategoryParams) error
}
