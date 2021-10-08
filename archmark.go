package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=view

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file. Expecting ENV to be set")
	}

	cfg := config.New()
	db := database.New(database.Config{
		ConnectionString: cfg.DBConnectionString(),
	})
	siteConfig, err := db.LoadSiteConfiguration()
	if err != nil || siteConfig == nil {
		panic(err)
	}

	r := gin.Default()
	router.Setup(r, cfg, db, *siteConfig)
	r.Run(":8080")
}
