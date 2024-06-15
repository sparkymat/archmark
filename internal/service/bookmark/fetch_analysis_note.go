package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) FetchAnalysisNote(ctx context.Context, id uuid.UUID) (dbx.AnalysisNote, error) {
	item, err := s.dbx.FetchAnalysisNoteByID(ctx, id)
	if err != nil {
		return dbx.AnalysisNote{}, fmt.Errorf("failed to fetch AnalysisNote: %w", err)
	}

	return item, nil
}
