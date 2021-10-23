package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/sparkymat/archmark/model"
)

func (s *service) ListAPITokens(ctx context.Context) ([]model.APIToken, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Select("*").
		From("api_tokens").
		OrderBy("created_at desc")

	querySQL, _, err := stmnt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to generate sql. err: %w", err)
	}

	log.Printf("SQL: %s\n", querySQL)

	var tokens []model.APIToken

	rows, err := s.conn.QueryxContext(ctx, querySQL)
	if err != nil {
		return nil, fmt.Errorf("failed to run query. err: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var token model.APIToken

		var deletedAt sql.NullTime

		err := rows.StructScan(&token)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row. err: %w", err)
		}

		if deletedAt.Valid {
			token.DeletedAt = &deletedAt.Time
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (s *service) DeleteAPIToken(ctx context.Context, id uint64) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Delete("api_tokens").
		Where(sq.Eq{"id": id})

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

func (s *service) CreateAPIToken(ctx context.Context, token string) (*model.APIToken, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Insert("api_tokens").
		Columns("token").
		Values(token).
		Suffix("RETURNING \"id\"")

	querySQL, args, err := stmnt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to generate sql. err: %w", err)
	}

	log.Printf("SQL: %s\n", querySQL)

	var id uint64

	err = s.conn.QueryRowxContext(ctx, querySQL, args...).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to run query. err: %w", err)
	}

	return &model.APIToken{
		ID:    id,
		Token: token,
	}, nil
}
