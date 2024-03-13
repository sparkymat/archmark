package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/view"
)

const SiteTitle = "archmark"

func Home(_ internal.Services) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := getCSRFToken(c)
		if csrfToken == "" {
			log.Print("error: csrf token not found")

			return c.String(http.StatusInternalServerError, "server error") //nolint:wrapcheck
		}

		// List js files in app

		files, err := os.ReadDir("./public/js/app")
		if err != nil {
			return c.String(http.StatusInternalServerError, "server error") //nolint:wrapcheck
		}

		filenames := lo.Map(files, func(file os.DirEntry, _ int) string {
			return file.Name()
		})

		jsFilenames := lo.Filter(filenames, func(filename string, _ int) bool { return filename[len(filename)-3:] == ".js" })
		cssFilenames := lo.Filter(filenames, func(filename string, _ int) bool { return filename[len(filename)-4:] == ".css" })

		pageHTML := view.Home()
		htmlString := view.BasicLayout(SiteTitle, csrfToken, cssFilenames, jsFilenames, pageHTML)

		return c.HTMLBlob(http.StatusOK, []byte(htmlString)) //nolint:wrapcheck
	}
}
