package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/internal/handler"
	"github.com/sparkymat/archmark/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID                uint           `gorm:"primaryKey"`
	CreatedAt         time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt         time.Time      `gorm:"default:current_timestamp"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	Username          string         `gorm:"index"`
	EncryptedPassword string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file. Expecting ENV to be set")
	}

	cfg := config.New()

	connectionString := fmt.Sprintf("host=localhost dbname=archmark port=5432 sslmode=disable")
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	db.AutoMigrate(&User{})

	r := gin.Default()
	r.Use(middleware.ConfigInjector(cfg))
	r.POST("/add", handler.Create)
	r.Run(":8080")
}
