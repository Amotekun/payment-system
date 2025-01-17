package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint   `gorm:"primaryKey"`
	Username         string `gorm:"uniqueIndex;not null"`
	Email            string `gorm:"uniqueIndex;not null"`
	PasswordHash     string `gorm:"not null"`
	Role             string `gorm:"type:varchar(20);default:'USER';not null"`
	TwoFactorEnabled bool   `gorm:"default:false"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
