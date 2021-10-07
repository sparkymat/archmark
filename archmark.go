package main

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

	db.AutoMigrate(
		&model.User{},
		&model.Bookmark{},
	)

	r := gin.Default()
	router.Setup(r, cfg, db)
	r.Run(":8080")
}
