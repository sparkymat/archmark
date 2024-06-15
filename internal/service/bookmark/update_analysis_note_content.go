package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) UpdateAnalysisNoteContent(ctx context.Context, id uuid.UUID, value string) (dbx.AnalysisNote, error) {
	input := dbx.UpdateAnalysisNoteContentParams{
		ID:      id,
		Content: value,
	}

	val, err := s.dbx.UpdateAnalysisNoteContent(ctx, input)
	if err != nil {
		return dbx.AnalysisNote{}, fmt.Errorf("failed to update AnalysisNote Content: %w", err)
	}

	return val, nil
}
