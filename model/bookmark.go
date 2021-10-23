package model

import (
	"time"
)

type Bookmark struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	URL       string
	Title     string
	Status    string
	Content   string
	FileName  string
}
