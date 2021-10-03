package config

import (
	"github.com/caarlos0/env/v6"
)

type API interface {
	ArchiveBoxPath() string
	ArchiveBoxUsername() string
	ArchiveBoxPassword() string
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

type envConfig struct {
	ArchiveBoxPath     string `env:"ARCHIVE_BOX_PATH,required"`
	ArchiveBoxUsername string `env:"ARCHIVE_BOX_USERNAME,required"`
	ArchiveBoxPassword string `env:"ARCHIVE_BOX_PASSWORD,required"`
}
