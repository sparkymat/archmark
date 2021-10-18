package presenter

import (
	"fmt"

	"github.com/sparkymat/archmark/model"
)

type Bookmark struct {
	Url         string
	OriginalUrl string
	Title       string
}

func PresentBookmark(bookmark model.Bookmark) Bookmark {
	return Bookmark{
		Url:         fmt.Sprintf("/b/%s", bookmark.FileName),
		OriginalUrl: bookmark.Url,
		Title:       bookmark.Title,
	}
}
