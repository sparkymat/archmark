package database

import (
	"errors"

	"github.com/sparkymat/archmark/model"
	"golang.org/x/crypto/bcrypt"
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
	LoadSiteConfiguration() (*model.Configuration, error)
	LoadAdminUser() (*model.User, error)
	CreateAdminUser(password string) (*model.User, error)
	FindUser(username string) (*model.User, error)
}

func New(cfg Config) API {
	conn, err := gorm.Open(postgres.Open(cfg.ConnectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(
		&model.User{},
		&model.Bookmark{},
		&model.Configuration{},
	)

	return &service{
		conn: conn,
	}
}

type service struct {
	conn *gorm.DB
}

func (s *service) LoadSiteConfiguration() (*model.Configuration, error) {
	var siteConfig model.Configuration
	result := s.conn.First(&siteConfig)
	if result.RowsAffected == 0 {
		result = s.conn.Create(&siteConfig)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return &siteConfig, nil
}

func (s *service) LoadAdminUser() (*model.User, error) {
	return s.FindUser("admin")
}

func (s *service) CreateAdminUser(password string) (*model.User, error) {
	encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), BcryptDefaultCost)
	if err != nil {
		return nil, err
	}
	admin := model.User{
		Username:          "admin",
		EncryptedPassword: string(encryptedPasswordBytes),
	}

	result := s.conn.Create(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}

func (s *service) FindUser(username string) (*model.User, error) {
	var user model.User

	result := s.conn.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 { // no results
		return nil, ErrNotFound
	}

	return &user, nil
}
