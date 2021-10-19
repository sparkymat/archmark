package model

import (
	"time"

	"gorm.io/gorm"
)

type ApiToken struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Token     string         `gorm:"not null;index"`
}