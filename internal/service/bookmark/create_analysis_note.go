package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/archmark/internal/dbx"
)

type CreateAnalysisNoteParams struct {
	Content    string
	BookmarkID uuid.UUID
}

func (s *Service) CreateAnalysisNote(ctx context.Context, params CreateAnalysisNoteParams) (dbx.AnalysisNote, error) {
	input := dbx.CreateAnalysisNoteParams{
		Content:    params.Content,
		Source:     dbx.AnalysisNoteSourceUser,
		BookmarkID: params.BookmarkID,
		State:      dbx.AnalysisNoteStateReady,
	}

	val, err := s.dbx.CreateAnalysisNote(ctx, input)
	if err != nil {
		return dbx.AnalysisNote{}, fmt.Errorf("failed to create AnalysisNote: %w", err)
	}

	return val, nil
}
