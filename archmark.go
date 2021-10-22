package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=view

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	r := echo.New()
	router.Setup(r, cfg, db)

	err = r.Start(":8080")
	if err != nil {
		panic(err)
	}
}
