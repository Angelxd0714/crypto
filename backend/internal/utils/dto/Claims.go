package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserId uuid.UUID `json:"user_id" binding:"required"`
	Email  string    `json:"email" binding:"required"`
	jwt.RegisteredClaims
}
