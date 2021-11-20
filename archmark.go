package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=view

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/archmark/app"
	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/database"
	"github.com/sparkymat/archmark/localize"
	"github.com/sparkymat/archmark/model"
	"github.com/sparkymat/archmark/router"
	"github.com/sparkymat/archmark/settings"
	"github.com/sparkymat/archmark/style"
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

	if err = db.AutoMigrate(); err != nil {
		panic(err)
	}

	localizer := localize.New()
	styler := style.New(style.LightTheme())

	settingsService, err := createSettingsService(context.Background(), cfg, db)
	if err != nil {
		panic(err)
	}

	appService := app.New(cfg, db, localizer, settingsService, styler)

	r := echo.New()
	router.Setup(r, appService)

	err = r.Start(":8080")
	if err != nil {
		panic(err)
	}
}

func createSettingsService(ctx context.Context, cfg *config.Service, db *database.Service) (*settings.Service, error) {
	settingsModel, err := db.LoadSettings(ctx, model.DefaultSettings(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to load settings. err: %w", err)
	}

	return settings.New(settingsModel, cfg), nil
}
