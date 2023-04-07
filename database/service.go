package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // importing postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // importing filesystem driver
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(connectionString string) (*Service, error) {
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pg conn. err: %w", err)
	}

	config.MinConns = 4
	config.MaxConns = 100

	dbConn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to pg. err: %w", err)
	}

	err = dbConn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to ping pg. err: %w", err)
	}

	return &Service{
		conn:             dbConn,
		connectionString: connectionString,
	}, nil
}

type Service struct {
	conn             *pgxpool.Pool
	connectionString string
}

func (s *Service) DB() *pgxpool.Pool {
	return s.conn
}

func (s *Service) AutoMigrate() error {
	m, err := migrate.New("file://./migrations", s.connectionString)
	if err != nil {
		return err //nolint:wrapcheck
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations. err: %w", err)
	}

	return nil
}
