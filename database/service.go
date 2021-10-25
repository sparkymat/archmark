package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// Importing file driver for migrations.
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

	// Bookmarks
	ListBookmarks(ctx context.Context, query string, page uint64, pageSize uint64) ([]model.Bookmark, error)
	CountBookmarks(ctx context.Context, query string) (uint64, error)
	FindBookmark(ctx context.Context, id uint64) (*model.Bookmark, error)
	CreateBookmark(ctx context.Context, bookmark *model.Bookmark) error
	MarkBookmarkCompleted(ctx context.Context, id uint64) error

	// API Tokens
	ListAPITokens(ctx context.Context) ([]model.APIToken, error)
	DeleteAPIToken(ctx context.Context, id uint64) error
	CreateAPIToken(ctx context.Context, token string) (*model.APIToken, error)
	LookupAPIToken(ctx context.Context, token string) (*model.APIToken, error)
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
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations. err: %w", err)
	}

	return nil
}
