package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/archive"
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

	var input CreateInput

	if c.Bind(&input) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}

	archiveAPI := archive.New()
	title, url, err := archiveAPI.Fetch(input.Url)
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
