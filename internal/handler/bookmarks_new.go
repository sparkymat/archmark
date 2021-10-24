package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/view"
)

func BookmarksNew(c echo.Context) error {
	csrfTokenVal := c.Get(middleware.DefaultCSRFConfig.ContextKey)
	if csrfTokenVal == nil {
		log.Print("error: csrf token not found")
		//nolint:wrapcheck
		return ShowError(c)
	}

	csrfToken, ok := csrfTokenVal.(string)
	if !ok {
		log.Print("error: csrf token not found")
		//nolint:wrapcheck
		return ShowError(c)
	}

	pageHTML := view.BookmarksNew(csrfToken)
	htmlString := view.Layout("archmark | add", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
