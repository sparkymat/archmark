package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
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

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
}
