package internal

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/service/bookmark"
)

type BookmarkService interface {
	CreateBookmark(ctx context.Context, params bookmark.CreateBookmarkParams) (dbx.Bookmark, error)
	SearchBookmarks(ctx context.Context, userID uuid.UUID, query string, pageSize int32, pageNumber int32) ([]dbx.Bookmark, int64, error)
	FetchRecentBookmarks(ctx context.Context, userID uuid.UUID, pageSize int32, pageNumber int32) ([]dbx.Bookmark, int64, error)
	FetchBookmark(ctx context.Context, id uuid.UUID) (dbx.Bookmark, error)
	DestroyBookmark(ctx context.Context, id uuid.UUID) error
	UpdateBookmarkTitle(ctx context.Context, id uuid.UUID, value string) (dbx.Bookmark, error)
	UpdateBookmarkTranscript(ctx context.Context, id uuid.UUID, value string) (dbx.Bookmark, error)
	UpdateBookmarkUrl(ctx context.Context, id uuid.UUID, value string) (dbx.Bookmark, error)
	UpdateBookmarkDescription(ctx context.Context, id uuid.UUID, value string) (dbx.Bookmark, error)
	UploadBookmarkImage(ctx context.Context, id uuid.UUID, filename string, attachmentFile io.Reader) (dbx.Bookmark, error)

	CreateAnalysisNote(ctx context.Context, params bookmark.CreateAnalysisNoteParams) (dbx.AnalysisNote, error)
	FetchRecentAnalysisNotes(ctx context.Context, parentID uuid.UUID) ([]dbx.AnalysisNote, int64, error)
	FetchAnalysisNote(ctx context.Context, id uuid.UUID) (dbx.AnalysisNote, error)
	DestroyAnalysisNote(ctx context.Context, id uuid.UUID) error
	UpdateAnalysisNoteContent(ctx context.Context, id uuid.UUID, value string) (dbx.AnalysisNote, error)
}
