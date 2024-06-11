package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) UpdateBookmarkTitle(ctx context.Context, id uuid.UUID, value string) (dbx.Bookmark, error) {
	input := dbx.UpdateBookmarkTitleParams{
		ID:    id,
		Title: value,
	}

	val, err := s.dbx.UpdateBookmarkTitle(ctx, input)
	if err != nil {
		return dbx.Bookmark{}, fmt.Errorf("failed to update Bookmark Title: %w", err)
	}

	return val, nil
}
