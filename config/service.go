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
	JWTSecret        string `env:"JWT_SECRET,required"`
	SessionSecret    string `env:"SESSION_SECRET,required"`
	DatabaseName     string `env:"DATABASE_NAME,required"`
	DatabaseHostname string `env:"DATABASE_HOSTNAME,required"`
	DatabasePort     string `env:"DATABASE_PORT,required"`
	DatabaseUsername string `env:"DATABASE_USERNAME"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	DatabaseSSLMode  bool   `env:"DATABASE_SSL_MODE" envDefault:"true"`
}

func (s *Service) JWTSecret() string {
	return s.envValues.JWTSecret
}

func (s *Service) SessionSecret() string {
	return s.envValues.SessionSecret
}

func (c *Service) DatabaseURL() string {
	connString := "postgres://"

	if c.envValues.DatabaseUsername != "" {
		connString = fmt.Sprintf("%s%s", connString, c.envValues.DatabaseUsername)

		if c.envValues.DatabasePassword != "" {
			encodedPassword := url.QueryEscape(c.envValues.DatabasePassword)
			connString = fmt.Sprintf("%s:%s", connString, encodedPassword)
		}

		connString = fmt.Sprintf("%s@", connString)
	}

	sslMode := "disable"
	if c.envValues.DatabaseSSLMode {
		sslMode = "require"
	}

	connString = fmt.Sprintf(
		"%s%s:%s/%s?sslmode=%s",
		connString,
		c.envValues.DatabaseHostname,
		c.envValues.DatabasePort,
		c.envValues.DatabaseName,
		sslMode,
	)

	return connString
}
