package model

import (
	"time"

	"gorm.io/gorm"
)

type Bookmark struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Url       string         `gorm:"not null"`
	Status    string         `gorm:"not null"`
	Content   string
}
