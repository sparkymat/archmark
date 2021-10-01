package main

import (
	"context"
	"log"
	"net/http"

	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
)

type AddInput struct {
	Url string `json:"url" binding:"required"`
}

func retrieve(url string) string {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", "1"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var body string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("body", &body, chromedp.ByQuery),
	); err != nil {
		log.Print(err)
		return ""
	}
	return body
}

func main() {
	r := gin.Default()
	r.POST("/add", func(c *gin.Context) {
		var input AddInput

		if c.Bind(&input) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "unimplemented",
			})
			return
		}

		log.Print(retrieve(input.Url))

		c.JSON(http.StatusOK, gin.H{
			"message": "got url",
		})
	})
	r.Run(":8080")
}
