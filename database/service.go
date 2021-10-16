package database

import (
	"errors"

	"github.com/sparkymat/archmark/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("not found")
)

const BcryptDefaultCost = 10

type Config struct {
	ConnectionString string
}

type API interface {
}

func New(cfg Config) API {
	conn, err := gorm.Open(postgres.Open(cfg.ConnectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(
		&model.Bookmark{},
	)

	return &service{
		conn: conn,
	}
}

type service struct {
	conn *gorm.DB
}
