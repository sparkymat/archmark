package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/auth"
	"github.com/sparkymat/archmark/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
)

//nolint:funlen
func BookmarksSearch(_ ConfigService, db DatabaseService) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, isUser := c.Get(auth.UserKey).(dbx.User)
		if !isUser {
			return renderError(c, http.StatusInternalServerError, "failed to load user", nil)
		}

		query := c.QueryParam("query")

		pageSizeString := c.QueryParam("page_size")

		pageSize, err := strconv.ParseInt(pageSizeString, 10, 32)
		if err != nil {
			return renderError(c, http.StatusBadRequest, "page_size was invalid", err)
		}

		pageNumberString := c.QueryParam("page_number")

		pageNumber, err := strconv.ParseInt(pageNumberString, 10, 32)
		if err != nil {
			return renderError(c, http.StatusBadRequest, "page_number was invalid", err)
		}

		offset := (pageNumber - 1) * pageSize

		var bookmarks []dbx.Bookmark
		var totalCount int64

		//nolint:nestif
		if query == "" {
			bookmarks, err = db.FetchBookmarksList(
				c.Request().Context(),
				dbx.FetchBookmarksListParams{
					UserID:     user.ID,
					PageOffset: int32(offset),
					PageLimit:  int32(pageSize),
				},
			)
		} else {
			bookmarks, err = db.SearchBookmarks(
				c.Request().Context(),
				dbx.SearchBookmarksParams{
					UserID:     user.ID,
					Query:      query,
					PageOffset: int32(offset),
					PageLimit:  int32(pageSize),
				},
			)
		}

		if err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to fetch bookmarks", err)
		}

		if query == "" {
			totalCount, err = db.CountBookmarksList(c.Request().Context(), user.ID)
		} else {
			totalCount, err = db.CountBookmarksSearchResults(c.Request().Context(), dbx.CountBookmarksSearchResultsParams{
				UserID: user.ID,
				Query:  query,
			})
		}

		if err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to fetch bookmarks count", err)
		}

		presentedBookmarks := []presenter.Bookmark{}

		for _, bookmark := range bookmarks {
			presentedBookmarks = append(presentedBookmarks, presenter.BookmarkFromModel(bookmark))
		}

		response := BookmarksListResponse{
			Items:      presentedBookmarks,
			PageSize:   int(pageSize),
			PageNumber: int(pageNumber),
			TotalCount: int(totalCount),
		}

		return c.JSON(http.StatusOK, response) //nolint:wrapcheck
	}
}
