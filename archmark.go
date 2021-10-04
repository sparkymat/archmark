package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sparkymat/archmark/archivebox"
	"github.com/sparkymat/archmark/config"
)

type AddInput struct {
	Url string `json:"url" binding:"required"`
}

const ConfigKey = "config"

func ConfigMiddleware(cfg config.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ConfigKey, cfg)
		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file. Expecting ENV to be set")
	}

	cfg := config.New()

	r := gin.Default()
	r.Use(ConfigMiddleware(cfg))
	r.POST("/add", func(c *gin.Context) {
		cfgVal, configFound := c.Get(ConfigKey)
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

		var input AddInput

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
	})
	r.Run(":8080")
}
