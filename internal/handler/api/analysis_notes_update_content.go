package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
)

type AnalysisNotesUpdateContentRequest struct {
	Content string `json:"content"`
}

func AnalysisNotesUpdateContent(s internal.Services) echo.HandlerFunc {
	return wrapWithAuthForMember(func(c echo.Context, _ dbx.User, id uuid.UUID) error {
		var request AnalysisNotesUpdateContentRequest
		if err := c.Bind(&request); err != nil {
			return renderError(c, http.StatusBadRequest, "invalid request", err)
		}

		item, err := s.Bookmark.UpdateAnalysisNoteContent(c.Request().Context(), id, request.Content)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "could not update AnalysisNote Content", err)
		}

		presentedAnalysisNote := presenter.AnalysisNoteFromModel(item)

		return c.JSON(http.StatusCreated, presentedAnalysisNote)
	})
}
