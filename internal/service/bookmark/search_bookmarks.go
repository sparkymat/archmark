package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) SearchBookmarks(ctx context.Context, userID uuid.UUID, query string, pageSize int32, pageNumber int32) ([]dbx.Bookmark, int64, error) {
	offset := (pageNumber - 1) * pageSize

	items, err := s.dbx.SearchBookmarks(ctx, dbx.SearchBookmarksParams{
		Query:      query,
		UserID:     userID,
		PageOffset: offset,
		PageLimit:  pageSize,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to search Bookmarks: %w", err)
	}

	totalCount, err := s.dbx.CountSearchedBookmarks(ctx, dbx.CountSearchedBookmarksParams{
		Query:  query,
		UserID: userID,
	})

	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch Bookmarks search count: %w", err)
	}

	return items, totalCount, nil
}
