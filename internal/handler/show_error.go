package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/view"
)

func ShowError(c echo.Context) error {
	return renderError(c, "Internal server error. Please try again later.")
}

func renderError(c echo.Context, message string) error {
	pageHTML := view.ShowError(message)
	htmlString := view.Layout("archmark", pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
