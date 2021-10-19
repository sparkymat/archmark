package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/presenter"
	"github.com/sparkymat/archmark/view"
)

const pageSize = 20

type HomeInput struct {
	Query string `query:"q"`
	Page  uint32 `query:"p"`
}

func Home(c echo.Context) error {
	input := &HomeInput{}
	if err := c.Bind(input); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if input.Page == 0 {
		input.Page = 1
	}

	dbVal := c.Get(middleware.DBKey)
	if dbVal == nil {
		return c.String(http.StatusInternalServerError, "db conn not found")
	}
	db := dbVal.(database.API)
	bookmarks, err := db.LoadBookmarks(input.Query, input.Page, pageSize)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	presentedBookmarks := []presenter.Bookmark{}
	for _, bookmark := range bookmarks {
		presentedBookmarks = append(presentedBookmarks, presenter.PresentBookmark(bookmark))
	}

	pageHTML := view.Home(presentedBookmarks)
	htmlString := view.Layout("archmark", pageHTML)
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
