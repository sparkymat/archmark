package config

import (
	"fmt"
	"net/url"

	"github.com/caarlos0/env/v7"
)

func New() (*Service, error) {
	var envValues envValues

	err := env.Parse(&envValues)
	if err != nil {
		return nil, fmt.Errorf("failed to load config from env. err: %w", err)
	}

	return &Service{
		envValues: envValues,
	}, nil
}

type Service struct {
	envValues envValues
}

type envValues struct {
	JWTSecret           string `env:"JWT_SECRET,required"`
	SessionSecret       string `env:"SESSION_SECRET,required"`
	DatabaseName        string `env:"DATABASE_NAME,required"`
	DatabaseHostname    string `env:"DATABASE_HOSTNAME,required"`
	DatabasePort        string `env:"DATABASE_PORT,required"`
	DatabaseUsername    string `env:"DATABASE_USERNAME"`
	DatabasePassword    string `env:"DATABASE_PASSWORD"`
	DatabaseSSLMode     bool   `env:"DATABASE_SSL_MODE" envDefault:"true"`
	DisableRegistration bool   `env:"DISABLE_REGISTRATION" envDefault:"false"`
	DownloadPath        string `env:"DOWNLOAD_PATH,required"`
	MonolithPath        string `env:"MONOLITH_PATH,required"`
}

func (s *Service) JWTSecret() string {
	return s.envValues.JWTSecret
}

func (s *Service) SessionSecret() string {
	return s.envValues.SessionSecret
}

func (s *Service) DatabaseURL() string {
	connString := "postgres://"

	if s.envValues.DatabaseUsername != "" {
		connString = fmt.Sprintf("%s%s", connString, s.envValues.DatabaseUsername)

		if s.envValues.DatabasePassword != "" {
			encodedPassword := url.QueryEscape(s.envValues.DatabasePassword)
			connString = fmt.Sprintf("%s:%s", connString, encodedPassword)
		}

		connString = fmt.Sprintf("%s@", connString)
	}

	sslMode := "disable"
	if s.envValues.DatabaseSSLMode {
		sslMode = "require"
	}

	connString = fmt.Sprintf(
		"%s%s:%s/%s?sslmode=%s",
		connString,
		s.envValues.DatabaseHostname,
		s.envValues.DatabasePort,
		s.envValues.DatabaseName,
		sslMode,
	)

	return connString
}

func (s *Service) DownloadPath() string {
	return s.envValues.DownloadPath
}

func (s *Service) MonolithPath() string {
	return s.envValues.MonolithPath
}

func (s *Service) DisableRegistration() bool {
	return s.envValues.DisableRegistration
}
