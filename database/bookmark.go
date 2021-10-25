package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/sparkymat/archmark/model"
)

//nolint:lll
func (s *service) ListBookmarks(ctx context.Context, query string, page uint64, pageSize uint64) ([]model.Bookmark, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Select("*").
		From("bookmarks")

	if query != "" {
		queryWords := strings.Split(query, " ")
		querySearchTerm := strings.Join(queryWords, "|")
		stmnt = stmnt.Where("weighted_tsv @@ to_tsquery(?)", querySearchTerm)
	} else {
		stmnt = stmnt.OrderBy("created_at desc")
	}

	offset := (page - 1) * pageSize
	stmnt = stmnt.Offset(offset).Limit(pageSize)

	querySQL, args, err := stmnt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to generate sql. err: %w", err)
	}

	log.Printf("SQL: %s\n", querySQL)

	var bookmarks []model.Bookmark

	rows, err := s.conn.QueryxContext(ctx, querySQL, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to run query. err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var bookmark model.Bookmark

		var deletedAt sql.NullTime

		err := rows.StructScan(&bookmark)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row. err: %w", err)
		}

		if deletedAt.Valid {
			bookmark.DeletedAt = &deletedAt.Time
		}

		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

func (s *service) FindBookmark(ctx context.Context, id uint64) (*model.Bookmark, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Select("*").
		From("bookmarks").
		Where("id = ?", id)

	querySQL, args, err := stmnt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to generate sql. err: %w", err)
	}

	log.Printf("SQL: %s\n", querySQL)

	var bookmark model.Bookmark

	err = s.conn.QueryRowxContext(ctx, querySQL, args...).StructScan(&bookmark)
	if err != nil {
		return nil, fmt.Errorf("failed to run query. err: %w", err)
	}

	return &bookmark, nil
}

func (s *service) CreateBookmark(ctx context.Context, bookmark *model.Bookmark) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Insert("bookmarks").
		Columns("url, title, status, content, file_name").
		Values(bookmark.URL, bookmark.Title, bookmark.Status, bookmark.Content, bookmark.FileName).
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

	bookmark.ID = id

	return nil
}

func (s *service) MarkBookmarkCompleted(ctx context.Context, id uint64) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Update("bookmarks").
		Set("status", "completed").
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

func (s *service) CountBookmarks(ctx context.Context, query string) (uint64, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Select("COUNT(*)").
		From("bookmarks")

	if query != "" {
		queryWords := strings.Split(query, " ")
		querySearchTerm := strings.Join(queryWords, "|")
		stmnt = stmnt.Where("weighted_tsv @@ to_tsquery(?)", querySearchTerm)
	}

	querySQL, args, err := stmnt.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to generate sql. err: %w", err)
	}

	log.Printf("SQL: %s\n", querySQL)

	var bookmarksCount uint64

	if err := s.conn.QueryRowxContext(ctx, querySQL, args...).Scan(&bookmarksCount); err != nil {
		return 0, fmt.Errorf("failed to run query. err: %w", err)
	}

	return bookmarksCount, nil
}
