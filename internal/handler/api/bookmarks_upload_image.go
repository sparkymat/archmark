package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
)

func BookmarksUploadImage(s internal.Services) echo.HandlerFunc {
	return wrapWithAuthForMember(func(c echo.Context, _ dbx.User, id uuid.UUID) error {
		fileHeader, err := c.FormFile("image_file")
		if err != nil {
			return renderError(c, http.StatusBadRequest, "missing file", err)
		}

		file, err := fileHeader.Open()
		if err != nil {
			return renderError(c, http.StatusBadRequest, "failed to open file", err)
		}
		defer file.Close()

		item, err := s.Bookmark.UploadBookmarkImage(
			c.Request().Context(),
			id,
			fileHeader.Filename,
			file,
		)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to upload image", err)
		}

		presentedBookmark := presenter.BookmarkFromModel(item)

		return c.JSON(http.StatusCreated, presentedBookmark)
	})
}
