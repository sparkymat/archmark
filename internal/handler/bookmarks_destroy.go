package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
	"github.com/sparkymat/archmark/archive"
)

func BookmarksDestroy(appService *app.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		bookmarkIDString := c.Param("id")

		bookmarkID, err := strconv.ParseUint(bookmarkIDString, base10, sixtyFourBits)
		if err != nil {
			log.Printf("error: %v", err)

			return ShowError(appService)(c)
		}

		bookmark, err := appService.DB.FindBookmark(c.Request().Context(), bookmarkID)
		if err != nil {
			log.Printf("error: %v", err)

			return ShowError(appService)(c)
		}

		if err = appService.DB.DeleteBookmark(c.Request().Context(), bookmarkID); err != nil {
			log.Printf("error: %v", err)

			return ShowError(appService)(c)
		}

		archiveAPI := archive.New(archive.Config{
			DownloadFolder: appService.Config.DownloadPath(),
		})

		if err = archiveAPI.RemoveArchiveFile(c.Request().Context(), bookmark.FileName); err != nil {
			log.Printf("error: %v", err)

			return ShowError(appService)(c)
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, c.Request().Referer())
	}
}
