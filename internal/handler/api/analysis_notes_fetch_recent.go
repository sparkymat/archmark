package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
)

type AnalysisNotesFetchRecentResponse struct {
	Items      []presenter.AnalysisNote `json:"items"`
	TotalCount int                      `json:"totalCount"`
}

func AnalysisNotesFetchRecent(s internal.Services) echo.HandlerFunc {
	return wrapWithAuthForChild(func(c echo.Context, _ dbx.User, parentID uuid.UUID) error {
		items, totalCount, err := s.Bookmark.FetchRecentAnalysisNotes(c.Request().Context(), parentID)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to fetch recent analysisNotes", err)
		}

		presentedItems := lo.Map(items, func(i dbx.AnalysisNote, _ int) presenter.AnalysisNote {
			return presenter.AnalysisNoteFromModel(i)
		})

		response := AnalysisNotesFetchRecentResponse{
			Items:      presentedItems,
			TotalCount: int(totalCount),
		}

		return c.JSON(http.StatusOK, response)
	})
}
