package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BookmarksArchive(_ ConfigService, db DatabaseService) echo.HandlerFunc {
	return func(c echo.Context) error {
		idString := c.Param("id")

		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err = db.ArchiveBookmark(c.Request().Context(), id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}
