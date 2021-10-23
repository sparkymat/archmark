package model

import (
	"time"
)

type APIToken struct {
	ID        uint64     `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Token     string     `db:"token"`
}

func (APIToken) TableName() string {
	return "api_tokens"
}
