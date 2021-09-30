package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddInput struct {
	Url string `json:"url" binding:"required"`
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

		c.JSON(http.StatusOK, gin.H{
			"message": "got url",
		})
	})
	r.Run(":8080")
}
