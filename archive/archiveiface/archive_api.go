package archiveiface

import (
	"context"

	"github.com/sparkymat/archmark/archive"
)

type ArchiveAPI interface {
	FetchDetails(ctx context.Context, url string, filename string) (*archive.ArchivedPage, error)
	RemoveArchiveFile(ctx context.Context, filename string) error
}
