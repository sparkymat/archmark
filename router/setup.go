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
	app := r.Group("")

	app.Use(middleware.ConfigInjector(cfg, db, siteConfig))
	app.Use(middleware.SetupRedirect(cfg, db))

	app.GET("/", handler.Home)
	app.GET("/setup", handler.Setup)
	app.POST("/add", handler.Create)

	r.Static("/css", "public/css")
	r.Static("/javascript", "public/javascript")
}
