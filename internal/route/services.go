package route

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sparkymat/archmark/dbx"
)

type ConfigService interface {
	JWTSecret() string
	SessionSecret() string
	DatabaseURL() string
	DisableRegistration() bool
	DownloadPath() string
	ReverseProxyAuthentication() bool
	ProxyAuthUsernameHeader() string
	ProxyAuthNameHeader() string
}

//nolint:interfacebloat
type DatabaseService interface {
	FetchUserByUsername(ctx context.Context, email string) (dbx.User, error)
	FetchBookmarksList(ctx context.Context, arg dbx.FetchBookmarksListParams) ([]dbx.Bookmark, error)
	CountBookmarksList(ctx context.Context, userID int64) (int64, error)
	CreateBookmark(ctx context.Context, arg dbx.CreateBookmarkParams) (dbx.Bookmark, error)
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
	SearchBookmarks(ctx context.Context, arg dbx.SearchBookmarksParams) ([]dbx.Bookmark, error)
	CountBookmarksSearchResults(ctx context.Context, arg dbx.CountBookmarksSearchResultsParams) (int64, error)
	FetchCategories(ctx context.Context, userID int64) ([]pgtype.Text, error)
	UpdateBookmarkCategory(ctx context.Context, arg dbx.UpdateBookmarkCategoryParams) error
	ArchiveBookmark(ctx context.Context, id int64) error
	UnarchiveBookmark(ctx context.Context, id int64) error
	FetchArchivedBookmarks(ctx context.Context, arg dbx.FetchArchivedBookmarksParams) ([]dbx.Bookmark, error)
}
