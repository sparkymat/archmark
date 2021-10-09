package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/internal/handler"
	"github.com/sparkymat/archmark/middleware"
	"github.com/sparkymat/archmark/model"
)

func Setup(r *gin.Engine, cfg config.API, db database.API, siteConfig model.Configuration) {
	app := r.Group("")

	app.Use(middleware.ConfigInjector(cfg, db, siteConfig))
	app.Use(middleware.SetupRedirect(cfg, db))

	app.GET("/", handler.Home)
	app.GET("/setup", handler.ShowSetup)
	app.POST("/setup", handler.DoSetup)
	app.POST("/add", handler.Create)

	r.Static("/css", "public/css")
	r.Static("/javascript", "public/javascript")
}
