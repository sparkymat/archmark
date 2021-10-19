package database

import (
	"errors"
	"log"
	"os"

	"github.com/sparkymat/archmark/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	CreateBookmark(bookmark *model.Bookmark) error
}

func New(cfg Config) API {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			Colorful: true,
		},
	)
	conn, err := gorm.Open(postgres.Open(cfg.ConnectionString), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(
		&model.Bookmark{},
		&model.ApiToken{},
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

	stmnt := s.conn

	if query != "" {
		stmnt = stmnt.Where("to_tsvector(content) @@ to_tsquery(?)", query)
	} else {
		stmnt = stmnt.Order("created_at desc")
	}

	if err := stmnt.Offset(offset).Limit(int(pageSize)).Find(&bookmarks); err.Error != nil {
		return nil, err.Error
	}
	return bookmarks, nil
}

func (s *service) CreateBookmark(bookmark *model.Bookmark) error {
	result := s.conn.Create(bookmark)
	return result.Error
}
