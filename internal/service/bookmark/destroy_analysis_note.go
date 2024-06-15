package bookmark

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"
)

func (s *Service) DestroyAnalysisNote(ctx context.Context, id uuid.UUID) error {
	folderPath := path.Join(s.storageFolder, "analysis_note", id.String())

	if err := os.RemoveAll(folderPath); err != nil {
		return fmt.Errorf("failed to remove AnalysisNote folder: %w", err)
	}

	err := s.dbx.DeleteAnalysisNote(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to fetch AnalysisNote: %w", err)
	}

	return nil
}
