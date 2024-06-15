package bookmark

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/sparkymat/archmark/internal/dbx"
)

type CreateAnalysisNoteParams struct {
	Content    string
	BookmarkID uuid.UUID
}

func (s *Service) CreateAnalysisNote(ctx context.Context, params CreateAnalysisNoteParams) (dbx.AnalysisNote, error) {
	// Fetch bookmark
	b, err := s.dbx.FetchBookmarkByID(ctx, params.BookmarkID)
	if err != nil {
		return dbx.AnalysisNote{}, fmt.Errorf("failed to fetch bookmark: %w", err)
	}

	// Existing conversation
	notes, err := s.dbx.FetchRecentAnalysisNotes(ctx, params.BookmarkID)
	if err != nil {
		return dbx.AnalysisNote{}, fmt.Errorf("failed to fetch recent AnalysisNotes: %w", err)
	}

	noteContents := lo.Map(notes, func(n dbx.AnalysisNote, _ int) string { return n.Content })

	input := dbx.CreateAnalysisNoteParams{
		Content:    params.Content,
		Source:     dbx.AnalysisNoteSourceUser,
		BookmarkID: params.BookmarkID,
		State:      dbx.AnalysisNoteStateReady,
	}

	createdNote, err := s.dbx.CreateAnalysisNote(ctx, input)
	if err != nil {
		return dbx.AnalysisNote{}, fmt.Errorf("failed to create question AnalysisNote: %w", err)
	}

	// Pass to llm
	answer, err := s.llm.AskQuestion(ctx, b.Transcript, noteContents, params.Content)
	if err != nil {
		return dbx.AnalysisNote{}, fmt.Errorf("failed to ask question to llm: %w", err)
	}

	input = dbx.CreateAnalysisNoteParams{
		Content:    answer,
		Source:     dbx.AnalysisNoteSourceService,
		BookmarkID: params.BookmarkID,
		State:      dbx.AnalysisNoteStateReady,
	}

	_, err = s.dbx.CreateAnalysisNote(ctx, input)
	if err != nil {
		return dbx.AnalysisNote{}, fmt.Errorf("failed to create response AnalysisNote: %w", err)
	}

	return createdNote, nil
}
