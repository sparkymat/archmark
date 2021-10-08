package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=view

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file. Expecting ENV to be set")
	}

	cfg := config.New()

	db, err := gorm.Open(postgres.Open(cfg.DBConnectionString()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&model.User{},
		&model.Bookmark{},
		&model.Configuration{},
	)

	var siteConfig model.Configuration
	result := db.First(&siteConfig)
	if result.RowsAffected == 0 {
		result = db.Create(&siteConfig)
		if result.Error != nil {
			panic(result.Error)
		}
	}

	r := gin.Default()
	router.Setup(r, cfg, db, siteConfig)
	r.Run(":8080")
}
