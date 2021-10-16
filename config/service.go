package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
)

type API interface {
	ArchiveBoxPath() string
	ArchiveBoxUsername() string
	ArchiveBoxPassword() string
	DBConnectionString() string
	AdminPassword() string
}

func New() API {
	envConfig := envConfig{}
	if err := env.Parse(&envConfig); err != nil {
		panic(err)
	}

	return &service{
		envConfig: envConfig,
	}
}

type service struct {
	envConfig envConfig
}

func (s *service) ArchiveBoxPath() string {
	return s.envConfig.ArchiveBoxPath
}

func (s *service) ArchiveBoxUsername() string {
	return s.envConfig.ArchiveBoxUsername
}

func (s *service) ArchiveBoxPassword() string {
	return s.envConfig.ArchiveBoxPassword
}

func (s *service) AdminPassword() string {
	return s.envConfig.AdminPassword
}

func (s *service) DBConnectionString() string {
	connFragments := []string{
		fmt.Sprintf("host=%s", s.envConfig.DBHostname),
		fmt.Sprintf("port=%d", s.envConfig.DBPort),
		fmt.Sprintf("dbname=%s", s.envConfig.DBDatabase),
	}

	if s.envConfig.DBSSLMode {
		connFragments = append(connFragments, "sslmode=require")
	} else {
		connFragments = append(connFragments, "sslmode=disable")
	}

	if s.envConfig.DBUsername != "" {
		connFragments = append(connFragments, fmt.Sprintf("user=%s", s.envConfig.DBUsername))
	}

	if s.envConfig.DBPassword != "" {
		connFragments = append(connFragments, fmt.Sprintf("password=%s", s.envConfig.DBPassword))
	}

	return strings.Join(connFragments, " ")
}

type envConfig struct {
	ArchiveBoxPath     string `env:"ARCHIVE_BOX_PATH,required"`
	ArchiveBoxUsername string `env:"ARCHIVE_BOX_USERNAME,required"`
	ArchiveBoxPassword string `env:"ARCHIVE_BOX_PASSWORD,required"`
	DBHostname         string `env:"DB_HOSTNAME,required"`
	DBPort             int64  `env:"DB_PORT,required"`
	DBUsername         string `env:"DB_USERNAME"`
	DBPassword         string `env:"DB_PASSWORD"`
	DBDatabase         string `env:"DB_DATABASE,required"`
	DBSSLMode          bool   `env:"DB_SSL_MODE"`
	AdminPassword      string `env:"ADMIN_PASSWORD,required"`
}
