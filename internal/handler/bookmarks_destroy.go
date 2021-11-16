package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/archive"
)

func BookmarksDestroy(c echo.Context) error {
	app := appServices(c)
	if app == nil {
		log.Print("error: app services not found")

		return ShowError(c)
	}

	bookmarkIDString := c.Param("id")

	bookmarkID, err := strconv.ParseUint(bookmarkIDString, base10, sixtyFourBits)
	if err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	bookmark, err := app.DB.FindBookmark(c.Request().Context(), bookmarkID)
	if err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	if err = app.DB.DeleteBookmark(c.Request().Context(), bookmarkID); err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	archiveAPI := archive.New(archive.Config{
		DownloadFolder: app.Config.DownloadPath(),
	})

	if err = archiveAPI.RemoveArchiveFile(c.Request().Context(), bookmark.FileName); err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, c.Request().Referer())
}
