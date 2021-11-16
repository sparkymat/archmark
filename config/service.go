package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/sparkymat/archmark/localize"
)

func New() *Service {
	envConfig := envConfig{}
	if err := env.Parse(&envConfig); err != nil {
		panic(err)
	}

	return &Service{
		envConfig: envConfig,
	}
}

type Service struct {
	envConfig envConfig
}

func (s *Service) AdminPassword() string {
	return s.envConfig.AdminPassword
}

func (s *Service) DBConnectionString() string {
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

func (s *Service) MonolithPath() string {
	return s.envConfig.MonolithPath
}

func (s *Service) DownloadPath() string {
	return s.envConfig.DownloadPath
}

func (s *Service) DefaultLanguage() localize.Language {
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
