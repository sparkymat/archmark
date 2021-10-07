package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/internal/handler"
	"github.com/sparkymat/archmark/middleware"
)

func Setup(r *gin.Engine, cfg config.API) {
	r.Use(middleware.ConfigInjector(cfg))

	r.POST("/add", handler.Create)
}
