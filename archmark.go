package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=view

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/dbx"
	"github.com/sparkymat/archmark/internal/jobs"
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

	// Initialize web server
	db := dbx.New(dbDriver.DB())
	route.Setup(e, appConfig, db)

	// Setup scheduler
	scheduler := gocron.NewScheduler(time.UTC)

	_, err = scheduler.Every(1).Hour().Do(jobs.DeleteExpiredBookmarks(appConfig, db))
	if err != nil {
		panic(err)
	}

	scheduler.StartAsync()

	e.Logger.Panic(e.Start(":8080"))
}
