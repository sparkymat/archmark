package database

import (
	"errors"
	"fmt"

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
	LoadBookmarks(query string, page uint32, pageSize uint32) ([]model.Bookmark, error)
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

	if result := s.conn.Find(bookmark, id); result.Error != nil {
		return nil, result.Error
	}

	return bookmark, nil
}

func (s *service) CreateBookmark(bookmark *model.Bookmark) error {
	result := s.conn.Create(bookmark)

	return result.Error
}

func (s *service) ListAPITokens() ([]model.APIToken, error) {
	var apiTokens []model.APIToken

	if result := s.conn.Find(&apiTokens); result.Error != nil {
		return nil, result.Error
	}

	return apiTokens, nil
}

func (s *service) DeleteAPIToken(id uint) error {
	err := s.conn.Delete(&model.APIToken{}, id)

	return err.Error
}

func (s *service) CreateAPIToken(token string) (*model.APIToken, error) {
	apiToken := &model.APIToken{
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

func (s *service) AutoMigrate() error {
	err := s.conn.AutoMigrate(
		&model.Bookmark{},
		&model.APIToken{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate. err: %w", err)
	}

	return nil
}
