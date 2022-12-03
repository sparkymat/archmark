package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file. Expecting ENV to be set")
	}

	r := echo.New()

	err = r.Start(":8080")
	if err != nil {
		panic(err)
	}
}
