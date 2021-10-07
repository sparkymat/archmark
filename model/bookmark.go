package model

import (
	"time"

	"gorm.io/gorm"
)

type Bookmark struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null"`
	User      User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Url       string         `gorm:"not null"`
	Status    string         `gorm:"not null"`
	Content   string
}
