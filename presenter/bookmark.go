package presenter

import (
	"fmt"
	"strings"

	"github.com/sparkymat/archmark/model"
	"github.com/xeonx/timeago"
)

type BookmarksList struct {
	Bookmarks    []Bookmark
	CurrentPage  uint64
	NextPageLink string
	ShowNextLink bool
}

type Bookmark struct {
	URL         string
	OriginalURL string
	IsActive    bool
	Title       string
	TimeSince   string
}

func PresentBookmarks(bookmarks []model.Bookmark, currentQuery string, currentPage uint64, hasMore bool) BookmarksList {
	presentedBookmarks := []Bookmark{}
	for _, bookmark := range bookmarks {
		presentedBookmarks = append(presentedBookmarks, PresentBookmark(bookmark))
	}

	queryFragments := []string{}
	if currentQuery != "" {
		queryFragments = append(queryFragments, fmt.Sprintf("q=%s", currentQuery))
	}
	if hasMore {
		queryFragments = append(queryFragments, fmt.Sprintf("p=%d", currentPage+1))
	}

	return BookmarksList{
		Bookmarks:    presentedBookmarks,
		CurrentPage:  currentPage,
		NextPageLink: fmt.Sprintf("/?%s", strings.Join(queryFragments, "&")),
		ShowNextLink: hasMore,
	}
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
