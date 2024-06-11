package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) FetchRecentBookmarks(ctx context.Context, userID uuid.UUID, pageSize int32, pageNumber int32) ([]dbx.Bookmark, int64, error) {
	offset := (pageNumber - 1) * pageSize

	items, err := s.dbx.FetchRecentBookmarks(ctx, dbx.FetchRecentBookmarksParams{
		PageOffset: offset,
		PageLimit:  pageSize,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get recent Bookmarks: %w", err)
	}

	totalCount, err := s.dbx.CountRecentBookmarks(ctx, userID)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch Bookmarks count: %w", err)
	}

	return items, totalCount, nil
}
