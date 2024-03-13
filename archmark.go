package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=internal/view

import (
	"log/slog"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/internal"
	"github.com/sparkymat/archmark/internal/config"
	"github.com/sparkymat/archmark/internal/database"
	"github.com/sparkymat/archmark/internal/dbx"
	"github.com/sparkymat/archmark/internal/route"
	"github.com/sparkymat/archmark/internal/service/user"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	dbDriver, err := database.New(cfg.DatabaseURL())
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	if err = dbDriver.AutoMigrate(); err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	// Initialize web server
	db := dbx.New(dbDriver.DB())

	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.RedisURL()})
	defer asynqClient.Close()

	userService := user.New(db)

	services := internal.Services{
		User: userService,
	}

	e := echo.New()
	route.Setup(e, cfg, services)

	e.Logger.Panic(e.Start(":8080"))
}
