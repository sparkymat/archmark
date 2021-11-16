package model

import (
	"time"

	"github.com/sparkymat/archmark/config"
)

type Settings struct {
	ID        uint64     `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Language  string     `db:"language"`
}

func DefaultSettings(cfg config.Service) Settings {
	return Settings{
		Language: string(cfg.DefaultLanguage()),
	}
}
