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
	Index       uint64
	ID          uint64
	URL         string
	OriginalURL string
	IsActive    bool
	Title       string
	TimeSince   string
}

//nolint:lll
func PresentBookmarks(bookmarks []model.Bookmark, currentQuery string, currentPage uint64, pageSize uint64, hasMore bool) BookmarksList {
	presentedBookmarks := []Bookmark{}
	startCount := (currentPage-1)*pageSize + 1

	for i, bookmark := range bookmarks {
		presentedBookmarks = append(presentedBookmarks, PresentBookmark(bookmark, startCount+uint64(i)))
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

func PresentBookmark(bookmark model.Bookmark, index uint64) Bookmark {
	return Bookmark{
		Index:       index,
		ID:          bookmark.ID,
		URL:         fmt.Sprintf("/b/%s", bookmark.FileName),
		OriginalURL: bookmark.URL,
		IsActive:    bookmark.Status == "completed",
		Title:       bookmark.Title,
		TimeSince:   timeago.English.Format(bookmark.CreatedAt),
	}
}
