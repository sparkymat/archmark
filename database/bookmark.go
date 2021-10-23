package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/sparkymat/archmark/model"
)

func (s *service) ListBookmarks(ctx context.Context, query string, page uint64, pageSize uint64) ([]model.Bookmark, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	stmnt := psql.
		Select("*").
		From("bookmarks")

	if query != "" {
		stmnt = stmnt.Where("to_tsvector(content) @@ to_tsquery(?)", query)
	} else {
		stmnt = stmnt.OrderBy("created_at desc")
	}

	offset := uint64((page - 1) * pageSize)
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

	var bookmarks model.Bookmark

	rows, err := s.conn.QueryxContext(ctx, querySQL, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to run query. err: %w", err)
	}

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
	panic("unimplemented")
	/*
		bookmark := &model.Bookmark{}

		if result := s.conn.Find(bookmark, id); result.Error != nil {
			return nil, result.Error
		}

		return bookmark, nil
	*/
}

func (s *service) CreateBookmark(ctx context.Context, bookmark *model.Bookmark) error {
	panic("unimplemented")
	/*
		result := s.conn.Create(bookmark)

		return result.Error
	*/
}

func (s *service) MarkBookmarkCompleted(ctx context.Context, id uint64) error {
	panic("unimplemented")
	/*
		result := s.conn.Model(&model.Bookmark{}).Where("id = ?", id).Update("status", "completed")

		return result.Error
	*/
}
