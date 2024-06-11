package bookmark

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"
)

func (s *Service) DestroyBookmark(ctx context.Context, id uuid.UUID) error {
	folderPath := path.Join(s.storageFolder, "bookmark", id.String())

	if err := os.RemoveAll(folderPath); err != nil {
		return fmt.Errorf("failed to remove Bookmark folder: %w", err)
	}

	err := s.dbx.DeleteBookmark(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to fetch Bookmark: %w", err)
	}

	return nil
}
