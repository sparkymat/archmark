//nolint:tagliatelle
package presenter

import (
	"fmt"
	"time"

	"github.com/sparkymat/archmark/dbx"
)

type Bookmark struct {
	ID        string  `json:"id"`
	Title     *string `json:"title"`
	Category  string  `json:"category"`
	URL       string  `json:"url"`
	FilePath  *string `json:"file_path"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func BookmarkFromModel(bookmark dbx.Bookmark) Bookmark {
	b := Bookmark{
		ID:        fmt.Sprintf("%d", bookmark.ID),
		URL:       bookmark.Url,
		Status:    string(bookmark.Status),
		CreatedAt: bookmark.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: bookmark.UpdatedAt.Time.Format(time.RFC3339),
	}

	if bookmark.Category.Valid {
		b.Category = bookmark.Category.String
	}

	if bookmark.Title.Valid {
		b.Title = &bookmark.Title.String
	}

	if bookmark.FilePath.Valid {
		b.FilePath = &bookmark.FilePath.String
	}

	return b
}
