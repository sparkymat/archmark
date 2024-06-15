package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) FetchRecentAnalysisNotes(ctx context.Context, parentID uuid.UUID) ([]dbx.AnalysisNote, int64, error) {
	items, err := s.dbx.FetchRecentAnalysisNotes(ctx, parentID)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get recent AnalysisNotes: %w", err)
	}

	totalCount, err := s.dbx.CountRecentAnalysisNotes(ctx, parentID)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch AnalysisNotes count: %w", err)
	}

	return items, totalCount, nil
}
