package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
)

func BookmarksDestroy(c echo.Context) error {
	dbVal := c.Get(middleware.DBKey)
	if dbVal == nil {
		log.Print("error: db conn not found")

		return ShowError(c)
	}

	db, ok := dbVal.(database.API)
	if !ok {
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
