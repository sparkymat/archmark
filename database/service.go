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
	FindBookmark(id uint) (*model.Bookmark, error)
	CreateBookmark(bookmark *model.Bookmark) error
	ListApiTokens() ([]model.ApiToken, error)
	DeleteApiToken(id uint) error
	CreateApiToken(token string) (*model.ApiToken, error)
	MarkBookmarkCompleted(id uint) error
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

	if result := stmnt.Offset(offset).Limit(int(pageSize)).Find(&bookmarks); result.Error != nil {
		return nil, result.Error
	}
	return bookmarks, nil
}

func (s *service) FindBookmark(id uint) (*model.Bookmark, error) {
	bookmark := &model.Bookmark{}
	result := s.conn.Find(bookmark, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return bookmark, nil
}

func (s *service) CreateBookmark(bookmark *model.Bookmark) error {
	result := s.conn.Create(bookmark)
	return result.Error
}

func (s *service) ListApiTokens() ([]model.ApiToken, error) {
	var apiTokens []model.ApiToken

	if result := s.conn.Find(&apiTokens); result.Error != nil {
		return nil, result.Error
	}
	return apiTokens, nil
}

func (s *service) DeleteApiToken(id uint) error {
	err := s.conn.Delete(&model.ApiToken{}, id)
	return err.Error
}

func (s *service) CreateApiToken(token string) (*model.ApiToken, error) {
	apiToken := &model.ApiToken{
		Token: token,
	}

	if result := s.conn.Create(&apiToken); result.Error != nil {
		return nil, result.Error
	}

	return apiToken, nil
}

func (s *service) MarkBookmarkCompleted(id uint) error {
	result := s.conn.Model(&model.Bookmark{}).Where("id = ?", id).Update("status", "completed")
	return result.Error
}
