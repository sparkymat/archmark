package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.String(http.StatusOK, "%d", time.Now().Unix())
}
