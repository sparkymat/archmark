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
	LoadBookmarks(query string, page uint32, pageSize uint32) ([]model.Bookmark, error)
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

func (s *service) LoadBookmarks(query string, page uint32, pageSize uint32) ([]model.Bookmark, error) {
	var bookmarks []model.Bookmark
	offset := int((page - 1) * pageSize)
	if err := s.conn.Where("title LIKE ?", query).Order("created_at desc").Offset(offset).Limit(int(pageSize)).Find(&bookmarks); err != nil {
		return nil, err.Error
	}
	return bookmarks, nil
}
