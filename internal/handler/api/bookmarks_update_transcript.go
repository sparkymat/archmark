package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
)

type BookmarksUpdateTranscriptRequest struct {
	Transcript string `json:"transcript"`
}

func BookmarksUpdateTranscript(s internal.Services) echo.HandlerFunc {
	return wrapWithAuthForMember(func(c echo.Context, _ dbx.User, id uuid.UUID) error {
		var request BookmarksUpdateTranscriptRequest
		if err := c.Bind(&request); err != nil {
			return renderError(c, http.StatusBadRequest, "invalid request", err)
		}

		item, err := s.Bookmark.UpdateBookmarkTranscript(c.Request().Context(), id, request.Transcript)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "could not update Bookmark Transcript", err)
		}

		presentedBookmark := presenter.BookmarkFromModel(item)

		return c.JSON(http.StatusCreated, presentedBookmark)
	})
}
