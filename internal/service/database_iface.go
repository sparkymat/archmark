package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

//nolint:interfacebloat
type DatabaseProvider interface {
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
	FetchUserByEmail(ctx context.Context, email string) (dbx.User, error)
	FetchUserByID(ctx context.Context, userID uuid.UUID) (dbx.User, error)

	CountRecentBookmarks(ctx context.Context, userID uuid.UUID) (int64, error)
	CountSearchedBookmarks(ctx context.Context, params dbx.CountSearchedBookmarksParams) (int64, error)
	CreateBookmark(ctx context.Context, params dbx.CreateBookmarkParams) (dbx.Bookmark, error)
	DeleteBookmark(ctx context.Context, id uuid.UUID) error
	FetchBookmarkByID(ctx context.Context, id uuid.UUID) (dbx.Bookmark, error)
	FetchBookmarksByIDs(ctx context.Context, ids []uuid.UUID) ([]dbx.Bookmark, error)
	FetchRecentBookmarks(ctx context.Context, arg dbx.FetchRecentBookmarksParams) ([]dbx.Bookmark, error)
	SearchBookmarks(ctx context.Context, arg dbx.SearchBookmarksParams) ([]dbx.Bookmark, error)
	UpdateBookmarkTitle(ctx context.Context, arg dbx.UpdateBookmarkTitleParams) (dbx.Bookmark, error)
	UpdateBookmarkTranscript(ctx context.Context, arg dbx.UpdateBookmarkTranscriptParams) (dbx.Bookmark, error)
	UpdateBookmarkUrl(ctx context.Context, arg dbx.UpdateBookmarkUrlParams) (dbx.Bookmark, error)
	UpdateBookmarkDescription(ctx context.Context, arg dbx.UpdateBookmarkDescriptionParams) (dbx.Bookmark, error)
	UpdateBookmarkImage(ctx context.Context, arg dbx.UpdateBookmarkImageParams) (dbx.Bookmark, error)

	CountRecentAnalysisNotes(ctx context.Context, parentID uuid.UUID) (int64, error)
	CreateAnalysisNote(ctx context.Context, params dbx.CreateAnalysisNoteParams) (dbx.AnalysisNote, error)
	DeleteAnalysisNote(ctx context.Context, id uuid.UUID) error
	FetchAnalysisNoteByID(ctx context.Context, id uuid.UUID) (dbx.AnalysisNote, error)
	FetchAnalysisNotesByIDs(ctx context.Context, ids []uuid.UUID) ([]dbx.AnalysisNote, error)
	FetchRecentAnalysisNotes(ctx context.Context, parentID uuid.UUID) ([]dbx.AnalysisNote, error)
	UpdateAnalysisNoteContent(ctx context.Context, arg dbx.UpdateAnalysisNoteContentParams) (dbx.AnalysisNote, error)
}
