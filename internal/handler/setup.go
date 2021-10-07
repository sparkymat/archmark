package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(c *gin.Context) {
	c.String(http.StatusOK, "setup")
}
