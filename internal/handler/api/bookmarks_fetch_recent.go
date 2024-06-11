package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
)

type BookmarksFetchRecentResponse struct {
	Items      []presenter.Bookmark `json:"items"`
	TotalCount int                  `json:"totalCount"`
	PageSize   int                  `json:"pageSize"`
	PageNumber int                  `json:"pageNumber"`
}

func BookmarksFetchRecent(s internal.Services) echo.HandlerFunc {
	return wrapWithAuth(func(c echo.Context, user dbx.User) error {

		pageSize, pageNumber, err := parsePaginationParams(c)
		if err != nil {
			return renderError(c, http.StatusBadRequest, "invalid pagination params", err)
		}

		items, totalCount, err := s.Bookmark.FetchRecentBookmarks(c.Request().Context(), user.ID, pageSize, pageNumber)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to fetch recent bookmarks", err)
		}

		presentedItems := lo.Map(items, func(i dbx.Bookmark, _ int) presenter.Bookmark {
			return presenter.BookmarkFromModel(i)
		})

		response := BookmarksFetchRecentResponse{
			Items:      presentedItems,
			PageSize:   int(pageSize),
			PageNumber: int(pageNumber),
			TotalCount: int(totalCount),
		}

		return c.JSON(http.StatusOK, response)
	})
}
