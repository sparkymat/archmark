package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
)

func BookmarksDestroy(s internal.Services) echo.HandlerFunc {
	return wrapWithAuthForMember(func(c echo.Context, _ dbx.User, id uuid.UUID) error {
		if err := s.Bookmark.DestroyBookmark(c.Request().Context(), id); err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to destroy bookmark", err)
		}

		return c.NoContent(http.StatusOK)
	})
}
