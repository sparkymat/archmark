package presenter

import "github.com/sparkymat/archmark/model"

type Bookmark struct {
	Url   string
	Title string
}

func PresentBookmark(bookmark model.Bookmark) Bookmark {
	return Bookmark{
		Url:   bookmark.Url,
		Title: bookmark.Title,
	}
}
