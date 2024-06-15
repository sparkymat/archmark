package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
	"github.com/sparkymat/archmark/internal/service/bookmark"
)

type AnalysisNotesCreateRequest struct {
	Content    string    `json:"content"`
	BookmarkId uuid.UUID `json:"bookmarkId"`
}

func AnalysisNotesCreate(s internal.Services) echo.HandlerFunc {
	return wrapWithAuth(func(c echo.Context, _ dbx.User) error {
		var request AnalysisNotesCreateRequest
		if err := c.Bind(&request); err != nil {
			return renderError(c, http.StatusBadRequest, "invalid request", err)
		}

		input := bookmark.CreateAnalysisNoteParams{}

		input.Content = request.Content
		input.BookmarkID = request.BookmarkId

		item, err := s.Bookmark.CreateAnalysisNote(
			c.Request().Context(),
			input,
		)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "could not create AnalysisNote", err)
		}

		presentedAnalysisNote := presenter.AnalysisNoteFromModel(item)

		return c.JSON(http.StatusCreated, presentedAnalysisNote)
	})
}
