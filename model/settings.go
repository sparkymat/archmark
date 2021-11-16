package model

import (
	"time"

	"github.com/sparkymat/archmark/config"
	"github.com/sparkymat/archmark/localize"
)

type Settings struct {
	ID        uint64     `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Language  string     `db:"language"`
}

func DefaultSettings(cfg *config.Service) Settings {
	if cfg == nil {
		return Settings{Language: string(localize.English)}
	}

	return Settings{Language: string(cfg.DefaultLanguage())}
}
