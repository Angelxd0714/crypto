package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Alert struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey"`
	UserID      uuid.UUID `gorm:"not null"`
	CryptoID    string    `gorm:"not null"`
	TargetPrice float64   `gorm:"not null"`
	Active      bool      `gorm:"not null;default:true"`
}
