package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
)

func AnalysisNotesShow(s internal.Services) echo.HandlerFunc {
	return wrapWithAuthForMember(func(c echo.Context, _ dbx.User, id uuid.UUID) error {
		item, err := s.Bookmark.FetchAnalysisNote(
			c.Request().Context(),
			id,
		)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to fetch analysisNote", err)
		}

		presentedItem := presenter.AnalysisNoteFromModel(item)

		return c.JSON(http.StatusOK, presentedItem)
	})
}
