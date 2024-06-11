package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
)

type BookmarksUpdateDescriptionRequest struct {
	Description string `json:"description"`
}

func BookmarksUpdateDescription(s internal.Services) echo.HandlerFunc {
	return wrapWithAuthForMember(func(c echo.Context, _ dbx.User, id uuid.UUID) error {
		var request BookmarksUpdateDescriptionRequest
		if err := c.Bind(&request); err != nil {
			return renderError(c, http.StatusBadRequest, "invalid request", err)
		}

		item, err := s.Bookmark.UpdateBookmarkDescription(c.Request().Context(), id, request.Description)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "could not update Bookmark Description", err)
		}

		presentedBookmark := presenter.BookmarkFromModel(item)

		return c.JSON(http.StatusCreated, presentedBookmark)
	})
}
