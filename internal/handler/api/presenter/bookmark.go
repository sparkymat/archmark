package presenter

import (
	"fmt"
	"time"

	"github.com/sparkymat/archmark/dbx"
)

type Bookmark struct {
	ID        string  `json:"id"`
	Title     *string `json:"title"`
	URL       string  `json:"url"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func BookmarkFromModel(bookmark dbx.Bookmark) Bookmark {
	b := Bookmark{
		ID:        fmt.Sprintf("%d", bookmark.ID),
		URL:       bookmark.Url,
		CreatedAt: bookmark.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: bookmark.UpdatedAt.Time.Format(time.RFC3339),
	}

	if bookmark.Title.Valid {
		b.Title = &bookmark.Title.String
	}

	return b
}
