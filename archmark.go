package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sparkymat/archmark/config"
)

type AddInput struct {
	Url string `json:"url" binding:"required"`
}

func retrieve(cfg config.API, url string) string {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", "1"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := context.WithTimeout(allocCtx, time.Second*5)

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("%s/admin/login/", cfg.ArchiveBoxPath())),
		chromedp.WaitVisible("#id_username", chromedp.ByQuery),
		chromedp.SendKeys("#id_username", cfg.ArchiveBoxUsername(), chromedp.ByQuery),
		chromedp.SendKeys("#id_password", cfg.ArchiveBoxPassword(), chromedp.ByQuery),
		chromedp.Submit("#id_username", chromedp.ByQuery),
		chromedp.WaitVisible("#recent-actions-module", chromedp.ByQuery),
		chromedp.FullScreenshot(&buf, 90),
	); err != nil {
		if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
			log.Print(err)
			return ""
		}
		log.Print(err)
		return ""
	}
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		log.Print(err)
		return ""
	}

	return ""
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
	log.Printf("%+v", cfg)

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

		var input AddInput

		if c.Bind(&input) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid input",
			})
			return
		}

		log.Print(retrieve(cfg, input.Url))

		c.JSON(http.StatusOK, gin.H{
			"message": "got url",
		})
	})
	r.Run(":8080")
}
