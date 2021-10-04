package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/archivebox"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/middleware"
)

type CreateInput struct {
	Url string `json:"url" binding:"required"`
}

func Create(c *gin.Context) {
	cfgVal, configFound := c.Get(middleware.ConfigKey)
	if !configFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "not configured",
		})
		return
	}

	cfg := cfgVal.(config.API)

	archiveBoxAPI := archivebox.New(archivebox.Config{
		Path:     cfg.ArchiveBoxPath(),
		Username: cfg.ArchiveBoxUsername(),
		Password: cfg.ArchiveBoxPassword(),
	})

	var input CreateInput

	if c.Bind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	title, url, err := archiveBoxAPI.ArchiveLink(input.Url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"title": title,
			"url":   url,
		})
	}
}
