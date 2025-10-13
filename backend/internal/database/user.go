package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                uuid.UUID `gorm:"primaryKey"`
	Username          string    `gorm:"unique;not null"`
	Email             string    `gorm:"unique;not null"`
	Password          string    `gorm:"not null"`
	SubscriptionLevel string    `gorm:"not null;default:'free'"` // e.g., free, premium
}
