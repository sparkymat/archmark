package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/sparkymat/archmark/localize"
)

type API interface {
	DBConnectionString() string
	AdminPassword() string
	MonolithPath() string
	DownloadPath() string
	DefaultLanguage() localize.Language
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

func (s *service) MonolithPath() string {
	return s.envConfig.MonolithPath
}

func (s *service) DownloadPath() string {
	return s.envConfig.DownloadPath
}

func (s *service) DefaultLanguage() localize.Language {
	return localize.LanguageFromString(s.envConfig.DefaultLanguage)
}

type envConfig struct {
	DBHostname      string `env:"DB_HOSTNAME,required"`
	DBPort          int64  `env:"DB_PORT,required"`
	DBUsername      string `env:"DB_USERNAME"`
	DBPassword      string `env:"DB_PASSWORD"`
	DBDatabase      string `env:"DB_DATABASE,required"`
	DBSSLMode       bool   `env:"DB_SSL_MODE"`
	AdminPassword   string `env:"ADMIN_PASSWORD,required"`
	MonolithPath    string `env:"MONOLITH_PATH,required"`
	DownloadPath    string `env:"DOWNLOAD_PATH,required"`
	DefaultLanguage string `env:"DEFAULT_LANGUAGE"`
}
