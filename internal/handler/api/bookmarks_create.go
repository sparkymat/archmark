package api

import (
	"net/http"

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

		presentedBookmark := presenter.BookmarkFromModel(bookmark)

		return c.JSON(http.StatusCreated, presentedBookmark)
	}
}
