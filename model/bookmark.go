package model

import (
	"time"
)

type Bookmark struct {
	ID           uint64     `db:"id"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
	URL          string     `db:"url"`
	Title        string     `db:"title"`
	Status       string     `db:"status"`
	Content      string     `db:"content"`
	FileName     string     `db:"file_name"`
	SearchVector string     `db:"weighted_tsv"`
}
