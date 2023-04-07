package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=view

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/dbx"
	"github.com/sparkymat/archmark/internal/route"
)

func main() {
	e := echo.New()

	appConfig, err := config.New()
	if err != nil {
		panic(err)
	}

	dbDriver, err := database.New(appConfig.DatabaseURL())
	if err != nil {
		log.Error(err)
		panic(err)
	}

	if err = dbDriver.AutoMigrate(); err != nil {
		log.Error(err)
		panic(err)
	}

	db := dbx.New(dbDriver.DB())

	route.Setup(e, appConfig, db)

	e.Logger.Panic(e.Start(":8080"))
}
