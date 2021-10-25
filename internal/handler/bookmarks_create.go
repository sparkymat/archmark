package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	faktory "github.com/contribsys/faktory/client"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/archive"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/model"
)

var ErrConfigNotFound = errors.New("config not found")

type BookmarksCreateInput struct {
	URL string `json:"url" form:"url" binding:"required"`
}

func APIBookmarksCreate(c echo.Context) error {
	bookmark, err := bookmarksCreate(c)
	if err != nil {
		//nolint:wrapcheck
		return c.JSON(http.StatusOK, map[string]string{
			"error": err.Error(),
		})
	}

	//nolint:wrapcheck
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": bookmark.ID,
	})
}

func BookmarksCreate(c echo.Context) error {
	if _, err := bookmarksCreate(c); err != nil {
		log.Printf("error: %v", err)

		return showError(c, "Unable to add bookmark. Please try again later.")
	}

	//nolint:wrapcheck
	return c.Redirect(http.StatusSeeOther, "/")
}

func bookmarksCreate(c echo.Context) (*model.Bookmark, error) {
	cfg := getConfig(c)
	db := getDB(c)

	if cfg == nil || db == nil {
		return nil, ErrConfigNotFound
	}

	var input BookmarksCreateInput

	if err := c.Bind(&input); err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	bookmark, err := createBookmark(c.Request().Context(), db, cfg, input.URL)
	if err != nil {
		return nil, err
	}

	err = queueDownloadJob(bookmark.ID)
	if err != nil {
		return nil, err
	}

	return bookmark, nil
}

func createBookmark(ctx context.Context, db database.API, cfg config.API, urlString string) (*model.Bookmark, error) {
	if _, err := url.ParseRequestURI(urlString); err != nil {
		return nil, fmt.Errorf("invalid url. err: %w", err)
	}

	archiveAPI := archive.New(archive.Config{
		DownloadFolder: cfg.DownloadPath(),
	})
	fileHash := strings.ReplaceAll(uuid.New().String(), "-", "")

	page, err := archiveAPI.Save(ctx, urlString, fileHash)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch page details. err: %w", err)
	}

	bookmark := model.Bookmark{
		URL:      urlString,
		Title:    page.Title,
		Status:   "pending",
		Content:  page.HTMLContent,
		FileName: fmt.Sprintf("%s.html", fileHash),
	}

	err = db.CreateBookmark(ctx, &bookmark)
	if err != nil {
		return nil, fmt.Errorf("failed to create bookmark in db. err: %w", err)
	}

	return &bookmark, nil
}

func queueDownloadJob(bookmarkID uint64) error {
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
