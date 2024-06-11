package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) FetchBookmark(ctx context.Context, id uuid.UUID) (dbx.Bookmark, error) {
	item, err := s.dbx.FetchBookmarkByID(ctx, id)
	if err != nil {
		return dbx.Bookmark{}, fmt.Errorf("failed to fetch Bookmark: %w", err)
	}

	return item, nil
}
