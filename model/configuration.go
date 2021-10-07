package model

import (
	"time"

	"gorm.io/gorm"
)

type Configuration struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	SiteName  string         `gorm:"default:archmark!"`
}
