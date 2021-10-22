package presenter

import (
	"fmt"

	"github.com/sparkymat/archmark/model"
	"github.com/xeonx/timeago"
)

type Bookmark struct {
	URL         string
	OriginalURL string
	IsActive    bool
	Title       string
	TimeSince   string
}

func PresentBookmark(bookmark model.Bookmark) Bookmark {
	return Bookmark{
		URL:         fmt.Sprintf("/b/%s", bookmark.FileName),
		OriginalURL: bookmark.URL,
		IsActive:    bookmark.Status == "completed",
		Title:       bookmark.Title,
		TimeSince:   timeago.English.Format(bookmark.CreatedAt),
	}
}
