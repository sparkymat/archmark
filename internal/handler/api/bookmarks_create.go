package api

import (
	"fmt"
	"net/http"

	faktory "github.com/contribsys/faktory/client"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/auth"
	"github.com/sparkymat/archmark/dbx"
	"github.com/sparkymat/archmark/internal/handler/api/presenter"
)

type BookmarksCreateRequest struct {
	URL string `json:"url"`
}

func BookmarksCreate(_ ConfigService, db DatabaseService) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, isUser := c.Get(auth.UserKey).(dbx.User)
		if !isUser {
			return renderError(c, http.StatusInternalServerError, "failed to load user", nil)
		}

		var request BookmarksCreateRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		bookmark, err := db.CreateBookmark(
			c.Request().Context(),
			dbx.CreateBookmarkParams{
				UserID: user.ID,
				Url:    request.URL,
			},
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		_ = queueDownloadJob(bookmark.ID)

		presentedBookmark := presenter.BookmarkFromModel(bookmark)

		return c.JSON(http.StatusCreated, presentedBookmark) //nolint:wrapcheck
	}
}

func queueDownloadJob(bookmarkID int64) error {
	client, err := faktory.Open()
	if err != nil {
		return fmt.Errorf("failed to connect to faktory. err: %w", err)
	}

	job := faktory.NewJob("SaveWebPage", fmt.Sprintf("%d", bookmarkID))

	err = client.Push(job)
	if err != nil {
		return fmt.Errorf("failed to queue job on Faktory. err: %w", err)
	}

	return nil
}
