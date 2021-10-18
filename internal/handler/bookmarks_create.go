package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/archive"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/model"
)

type BookmarksCreateInput struct {
	Url string `json:"url" form:"url" binding:"required"`
}

func BookmarksCreate(c echo.Context) error {
	cfgVal := c.Get(middleware.ConfigKey)
	dbVal := c.Get(middleware.DBKey)
	if cfgVal == nil || dbVal == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "not configured",
		})
	}
	cfg := cfgVal.(config.API)
	db := dbVal.(database.API)

	var input BookmarksCreateInput

	if c.Bind(&input) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	archiveAPI := archive.New(archive.Config{
		MonolithPath:   cfg.MonolithPath(),
		DownloadFolder: cfg.DownloadPath(),
	})
	fileHash := strings.ReplaceAll(uuid.New().String(), "-", "")
	page, err := archiveAPI.Save(input.Url, fileHash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	bookmark := model.Bookmark{
		Url:      input.Url,
		Title:    page.Title,
		Status:   "pending",
		Content:  page.HTMLContent,
		FileName: fmt.Sprintf("%s.html", fileHash),
	}

	err = db.CreateBookmark(&bookmark)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
