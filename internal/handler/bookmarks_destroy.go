package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BookmarksDestroy(c echo.Context) error {
	db := getDB(c)
	if db == nil {
		log.Print("error: db conn not found")

		return ShowError(c)
	}

	bookmarkIDString := c.Param("id")

	bookmarkID, err := strconv.ParseUint(bookmarkIDString, base10, sixtyFourBits)
	if err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	if err = db.DeleteBookmark(c.Request().Context(), bookmarkID); err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, c.Request().Referer())
}
