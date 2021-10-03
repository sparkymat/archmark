package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sparkymat/archmark/config"
)

type AddInput struct {
	Url string `json:"url" binding:"required"`
}

func retrieve(cfg config.API, url string) (string, string, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", "1"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := context.WithTimeout(allocCtx, time.Second*30)

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var title, fileUrl string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("%s/admin/login/", cfg.ArchiveBoxPath())),
		chromedp.WaitVisible("#id_username", chromedp.ByQuery),
		chromedp.SendKeys("#id_username", cfg.ArchiveBoxUsername(), chromedp.ByQuery),
		chromedp.SendKeys("#id_password", cfg.ArchiveBoxPassword(), chromedp.ByQuery),
		chromedp.Submit("#id_username", chromedp.ByQuery),
		chromedp.WaitVisible("#recent-actions-module", chromedp.ByQuery),
		chromedp.Navigate(fmt.Sprintf("%s/add/", cfg.ArchiveBoxPath())),
		chromedp.WaitVisible("#id_url", chromedp.ByQuery),
		chromedp.SendKeys("#id_url", url, chromedp.ByQuery),
		chromedp.Click("#id_depth_0", chromedp.ByQuery),
		chromedp.QueryAfter("#id_archive_methods > option[value=\"title\"]", func(ctx context.Context, id runtime.ExecutionContextID, nodes ...*cdp.Node) error {
			return chromedp.MouseClickNode(nodes[0], chromedp.ButtonModifiers(input.ModifierCtrl)).Do(ctx)
		}, chromedp.ByQuery),
		chromedp.QueryAfter("#id_archive_methods > option[value=\"singlefile\"]", func(ctx context.Context, id runtime.ExecutionContextID, nodes ...*cdp.Node) error {
			return chromedp.MouseClickNode(nodes[0], chromedp.ButtonModifiers(input.ModifierCtrl)).Do(ctx)
		}, chromedp.ByQuery),
		chromedp.Click("#submit", chromedp.ByQuery),
		chromedp.WaitVisible("pre#stdout", chromedp.ByQuery),
		chromedp.Navigate(fmt.Sprintf("%s/public/", cfg.ArchiveBoxPath())),
		chromedp.WaitVisible("#searchbar", chromedp.ByQuery),
		chromedp.SendKeys("#searchbar", url, chromedp.ByQuery),
		chromedp.SendKeys("#searchbar", kb.Enter, chromedp.ByQuery),
		chromedp.WaitVisible("#searchbar", chromedp.ByQuery),
		chromedp.Text(".title-col > a:nth-child(2) > span:nth-child(1)", &title, chromedp.ByQuery),
		chromedp.AttributeValue("a[title='singlefile']", "href", &fileUrl, nil, chromedp.ByQuery),
	); err != nil {
		return "", "", err
	}

	return title, fmt.Sprintf("%s%s", cfg.ArchiveBoxPath(), fileUrl), nil
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

		title, url, err := retrieve(cfg, input.Url)

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
