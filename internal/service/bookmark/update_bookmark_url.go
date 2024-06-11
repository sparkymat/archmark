package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) UpdateBookmarkUrl(ctx context.Context, id uuid.UUID, value string) (dbx.Bookmark, error) {
	input := dbx.UpdateBookmarkUrlParams{
		ID:  id,
		Url: value,
	}

	val, err := s.dbx.UpdateBookmarkUrl(ctx, input)
	if err != nil {
		return dbx.Bookmark{}, fmt.Errorf("failed to update Bookmark Url: %w", err)
	}

	return val, nil
}
