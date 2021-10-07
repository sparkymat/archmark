package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/internal/handler"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/model"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, cfg config.API, db *gorm.DB, siteConfig model.Configuration) {
	r.Use(middleware.ConfigInjector(cfg, db, siteConfig))
	r.Use(middleware.SetupRedirect(cfg, db))

	r.GET("/", handler.Home)
	r.GET("/setup", handler.Setup)
	r.POST("/add", handler.Create)
}
