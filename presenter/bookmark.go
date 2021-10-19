package presenter

import (
	"fmt"

	"github.com/sparkymat/archmark/model"
	"github.com/xeonx/timeago"
)

type Bookmark struct {
	Url         string
	OriginalUrl string
	Title       string
	TimeSince   string
}

func PresentBookmark(bookmark model.Bookmark) Bookmark {
	return Bookmark{
		Url:         fmt.Sprintf("/b/%s", bookmark.FileName),
		OriginalUrl: bookmark.Url,
		Title:       bookmark.Title,
		TimeSince:   timeago.English.Format(bookmark.CreatedAt),
	}
}
