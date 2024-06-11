package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

type CreateBookmarkParams struct {
	Title       string
	Url         string
	Description string

	UserId uuid.UUID
}

func (s *Service) CreateBookmark(ctx context.Context, params CreateBookmarkParams) (dbx.Bookmark, error) {
	input := dbx.CreateBookmarkParams{
		Title:       params.Title,
		Url:         params.Url,
		Description: params.Description,

		UserID: params.UserId,
	}

	val, err := s.dbx.CreateBookmark(ctx, input)
	if err != nil {
		return dbx.Bookmark{}, fmt.Errorf("failed to create Bookmark: %w", err)
	}

	return val, nil
}
