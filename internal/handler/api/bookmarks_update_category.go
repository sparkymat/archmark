package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/dbx"
)

type BookmarksUpdateCategoryRequest struct {
	Category string `json:"category"`
}

func BookmarksUpdateCategory(_ ConfigService, db DatabaseService) echo.HandlerFunc {
	return func(c echo.Context) error {
		idString := c.Param("id")

		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var request BookmarksUpdateCategoryRequest
		if err = c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err = db.UpdateBookmarkCategory(c.Request().Context(), dbx.UpdateBookmarkCategoryParams{
			ID:       id,
			Category: request.Category,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}
