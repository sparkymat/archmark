package bookmark

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sparkymat/archmark/internal/dbx"
)

func (s *Service) UploadBookmarkImage(ctx context.Context, id uuid.UUID, filename string, attachmentFile io.Reader) (dbx.Bookmark, error) {
	folderPath := path.Join(s.storageFolder, "bookmark", id.String())

	if err := os.MkdirAll(folderPath, 0o755); err != nil { //nolint:gomnd
		return dbx.Bookmark{}, fmt.Errorf("failed to create bookmark folder. err: %w", err)
	}

	filePath := path.Join(folderPath, filename)

	data, err := io.ReadAll(attachmentFile)
	if err != nil {
		return dbx.Bookmark{}, fmt.Errorf("failed to read attachment file: %w", err)
	}

	//nolint:gomnd
	if err = os.WriteFile(filePath, data, 0o600); err != nil {
		return dbx.Bookmark{}, fmt.Errorf("failed to write attachment file: %w", err)
	}

	input := dbx.UpdateBookmarkImageParams{
		ID:    id,
		Image: pgtype.Text{String: fmt.Sprintf("/bookmark/%s/%s", id.String(), filename), Valid: true},
	}

	item, err := s.dbx.UpdateBookmarkImage(ctx, input)
	if err != nil {
		return dbx.Bookmark{}, fmt.Errorf("failed to update artist photo: %w", err)
	}

	return item, nil
}
