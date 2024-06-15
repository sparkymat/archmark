package config

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/caarlos0/env/v11"
	"github.com/sparkymat/archmark/internal/providers/llm"
)

var ErrInvalidLLMEngine = errors.New("invalid LLM engine")

func New() (*Service, error) {
	var envValues envValues

	err := env.Parse(&envValues)
	if err != nil {
		return nil, fmt.Errorf("failed to load config from env. err: %w", err)
	}

	s := &Service{
		envValues: envValues,
	}

	switch envValues.LLMEngine {
	case string(llm.EngineTypeGPT3):
		s.llmEngine = llm.EngineTypeGPT3
		s.llmKey = envValues.LLMKey
	default:
		return nil, fmt.Errorf("invalid LLM engine %s: %w", envValues.LLMEngine, ErrInvalidLLMEngine)
	}

	return s, nil
}

type Service struct {
	envValues envValues
	llmEngine llm.EngineType
	llmKey    string
}

type envValues struct {
	JWTSecret                  string `env:"JWT_SECRET,required"`
	SessionSecret              string `env:"SESSION_SECRET,required"`
	DatabaseName               string `env:"DATABASE_NAME,required"`
	DatabaseHostname           string `env:"DATABASE_HOSTNAME,required"`
	DatabasePort               string `env:"DATABASE_PORT,required"`
	DatabaseUsername           string `env:"DATABASE_USERNAME"`
	DatabasePassword           string `env:"DATABASE_PASSWORD"`
	DatabaseSSLMode            bool   `env:"DATABASE_SSL_MODE"            envDefault:"true"`
	DisableRegistration        bool   `env:"DISABLE_REGISTRATION"         envDefault:"false"`
	ReverseProxyAuthentication bool   `env:"REVERSE_PROXY_AUTHENTICATION" envDefault:"false"`
	ProxyAuthEmailHeader       string `env:"PROXY_AUTH_EMAIL_HEADER"      envDefault:"Remote-Email"`
	ProxyAuthNameHeader        string `env:"PROXY_AUTH_NAME_HEADER"       envDefault:"Remote-Name"`
	RedisURL                   string `env:"REDIS_URL,required"`
	StorageFolder              string `env:"STORAGE_FOLDER,required"`
	LLMEngine                  string `env:"LLM_ENGINE,required"`
	LLMKey                     string `env:"LLM_KEY"`
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

		connString += "@"
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

func (s *Service) DisableRegistration() bool {
	return s.envValues.DisableRegistration
}

func (s *Service) ReverseProxyAuthentication() bool {
	return s.envValues.ReverseProxyAuthentication
}

func (s *Service) ProxyAuthEmailHeader() string {
	return s.envValues.ProxyAuthEmailHeader
}

func (s *Service) ProxyAuthNameHeader() string {
	return s.envValues.ProxyAuthNameHeader
}

func (s *Service) RedisURL() string {
	return s.envValues.RedisURL
}

func (s *Service) StorageFolder() string {
	return s.envValues.StorageFolder
}

func (s *Service) LLMEngine() llm.EngineType {
	return s.llmEngine
}

func (s *Service) LLMKey() string {
	return s.llmKey
}
