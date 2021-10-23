package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"github.com/sparkymat/archmark/model"
)

var ErrNotFound = errors.New("not found")

const BcryptDefaultCost = 10

type Config struct {
	ConnectionString string
}

type API interface {
	AutoMigrate() error
	ListBookmarks(query string, page uint32, pageSize uint32) ([]model.Bookmark, error)
	FindBookmark(id uint) (*model.Bookmark, error)
	CreateBookmark(bookmark *model.Bookmark) error
	ListAPITokens() ([]model.APIToken, error)
	DeleteAPIToken(id uint) error
	CreateAPIToken(token string) (*model.APIToken, error)
	MarkBookmarkCompleted(id uint) error
}

func New(cfg Config) API {
	dbConn, err := sqlx.Connect("postgres", cfg.ConnectionString)
	if err != nil {
		panic(err)
	}

	err = dbConn.Ping()
	if err != nil {
		panic(err)
	}

	return &service{
		conn: dbConn,
	}
}

type service struct {
	conn *sqlx.DB
}

func (s *service) AutoMigrate() error {
	driver, err := postgres.WithInstance(s.conn.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver. err: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration driver. err: %w", err)
	}

	err = m.Up()
	if err != nil {
		return fmt.Errorf("failed to apply migrations. err: %w", err)
	}

	return nil
}
