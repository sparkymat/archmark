package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                uint           `gorm:"primaryKey"`
	CreatedAt         time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt         time.Time      `gorm:"default:current_timestamp"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	Username          string         `gorm:"not null;index"`
	EncryptedPassword string         `gorm:"not null"`
}
