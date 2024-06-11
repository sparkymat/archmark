package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) UpdateBookmarkDescription(ctx context.Context, id uuid.UUID, value string) (dbx.Bookmark, error) {
	input := dbx.UpdateBookmarkDescriptionParams{
		ID:          id,
		Description: value,
	}

	val, err := s.dbx.UpdateBookmarkDescription(ctx, input)
	if err != nil {
		return dbx.Bookmark{}, fmt.Errorf("failed to update Bookmark Description: %w", err)
	}

	return val, nil
}
