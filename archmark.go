package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "unimplemented",
		})
	})
	r.Run(":8080")
}
