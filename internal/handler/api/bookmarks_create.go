package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
	"github.com/sparkymat/archmark/internal/service/bookmark"
)

type BookmarksCreateRequest struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func BookmarksCreate(s internal.Services) echo.HandlerFunc {
	return wrapWithAuth(func(c echo.Context, user dbx.User) error {
		var request BookmarksCreateRequest
		if err := c.Bind(&request); err != nil {
			return renderError(c, http.StatusBadRequest, "invalid request", err)
		}

		input := bookmark.CreateBookmarkParams{}

		input.Title = request.Title
		input.Url = request.Url
		input.Description = request.Description
		input.UserId = user.ID

		item, err := s.Bookmark.CreateBookmark(
			c.Request().Context(),
			input,
		)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "could not create Bookmark", err)
		}

		presentedBookmark := presenter.BookmarkFromModel(item)

		return c.JSON(http.StatusCreated, presentedBookmark)
	})
}
