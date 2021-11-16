package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/sparkymat/archmark/model"
)

func (s *service) LoadSettings(ctx context.Context, defaultSettings model.Settings) (*model.Settings, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Select("*").
		From("settings").
		OrderBy("id desc").
		Limit(1)

	querySQL, _, err := stmnt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to generate sql. err: %w", err)
	}

	log.Printf("SQL: %s\n", querySQL)

	var settings model.Settings

	err = s.conn.QueryRowxContext(ctx, querySQL).StructScan(&settings)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed to run query. err: %w", err)
		}

		copiedSettings := defaultSettings

		err = s.createSettings(ctx, &copiedSettings)
		if err != nil {
			return nil, fmt.Errorf("failed to create settings. err: %w", err)
		}

		settings = copiedSettings
	}

	return &settings, nil
}

func (s *service) UpdateSettings(ctx context.Context, settings *model.Settings) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Insert("settings").
		Columns("language").
		Values(settings.Language)

	querySQL, args, err := stmnt.ToSql()
	if err != nil {
		return fmt.Errorf("failed to generate sql. err: %w", err)
	}

	log.Printf("SQL: %s\n", querySQL)

	_, err = s.conn.ExecContext(ctx, querySQL, args...)
	if err != nil {
		return fmt.Errorf("failed to run query. err: %w", err)
	}

	return nil
}

func (s *service) createSettings(ctx context.Context, settings *model.Settings) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Insert("settings").
		Columns("language").
		Values(settings.Language).
		Suffix("RETURNING \"id\"")

	querySQL, args, err := stmnt.ToSql()
	if err != nil {
		return fmt.Errorf("failed to generate sql. err: %w", err)
	}

	log.Printf("SQL: %s\n", querySQL)

	var id uint64

	err = s.conn.QueryRowxContext(ctx, querySQL, args...).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to run query. err: %w", err)
	}

	settings.ID = id

	return nil
}
