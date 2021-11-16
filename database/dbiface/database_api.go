package dbiface

import (
	"context"

	"github.com/sparkymat/archmark/model"
)

type DatabaseAPI interface {
	AutoMigrate() error

	// Bookmarks
	ListBookmarks(ctx context.Context, query string, page uint64, pageSize uint64) ([]model.Bookmark, error)
	CountBookmarks(ctx context.Context, query string) (uint64, error)
	FindBookmark(ctx context.Context, id uint64) (*model.Bookmark, error)
	CreateBookmark(ctx context.Context, bookmark *model.Bookmark) error
	MarkBookmarkCompleted(ctx context.Context, id uint64) error
	DeleteBookmark(ctx context.Context, id uint64) error

	// API Tokens
	ListAPITokens(ctx context.Context) ([]model.APIToken, error)
	DeleteAPIToken(ctx context.Context, id uint64) error
	CreateAPIToken(ctx context.Context, token string) (*model.APIToken, error)
	LookupAPIToken(ctx context.Context, token string) (*model.APIToken, error)

	// Settings
	LoadSettings(ctx context.Context, defaultSettings model.Settings) (*model.Settings, error)
	UpdateSettings(ctx context.Context, settings *model.Settings) error
}
