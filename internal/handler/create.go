package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/archivebox"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/middleware"
)

type CreateInput struct {
	Url string `json:"url" binding:"required"`
}

func Create(c echo.Context) error {
	cfgVal := c.Get(middleware.ConfigKey)
	if cfgVal == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "not configured",
		})
	}

	cfg := cfgVal.(config.API)

	archiveBoxAPI := archivebox.New(archivebox.Config{
		Path:     cfg.ArchiveBoxPath(),
		Username: cfg.ArchiveBoxUsername(),
		Password: cfg.ArchiveBoxPassword(),
	})

	var input CreateInput

	if c.Bind(&input) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	title, url, err := archiveBoxAPI.ArchiveLink(input.Url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"title": title,
		"url":   url,
	})
}
