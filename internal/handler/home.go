package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/presenter"
	"github.com/sparkymat/archmark/view"
)

const pageSize = 20

type HomeInput struct {
	Query string `query:"q"`
	Page  uint64 `query:"p"`
}

func Home(c echo.Context) error {
	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		return ShowError(c)
	}

	app := appServices(c)
	if app == nil {
		log.Print("error: app services not found")

		return ShowError(c)
	}

	input := &HomeInput{}
	if err := c.Bind(input); err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	if input.Page == 0 {
		input.Page = 1
	}

	bookmarks, err := app.DB.ListBookmarks(c.Request().Context(), input.Query, input.Page, pageSize)
	if err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	bookmarksCount, err := app.DB.CountBookmarks(c.Request().Context(), input.Query)
	if err != nil {
		log.Printf("error: %v", err)

		return ShowError(c)
	}

	presentedBookmarks := presenter.PresentBookmarks(bookmarks, input.Query, input.Page, pageSize, bookmarksCount)
	pageHTML := view.Home(*app.Localizer, app.Settings.Language(), csrfToken, input.Query != "", input.Query, presentedBookmarks)
	htmlString := view.Layout(*app.Localizer, app.Settings.Language(), "archmark", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
