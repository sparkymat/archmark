package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/archmark/view"
)

func BookmarksNew(c echo.Context) error {
	csrfTokenVal := c.Get(middleware.DefaultCSRFConfig.ContextKey)
	if csrfTokenVal == nil {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "csrf token not found")
	}

	csrfToken, ok := csrfTokenVal.(string)
	if !ok {
		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "csrf token not found")
	}

	pageHTML := view.BookmarksNew(csrfToken)
	htmlString := view.Layout("archmark | add", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
