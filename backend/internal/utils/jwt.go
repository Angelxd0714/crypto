package utils

import (
	"gin-quickstart/internal/database"
	"gin-quickstart/internal/utils/dto"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GetJWTSecret() []byte {
	if len(jwtSecret) == 0 {
		jwtSecret = []byte("default_secret_key")
	}
	return jwtSecret
}

func GenerateJWTSecret(user *database.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &dto.Claims{
		UserId: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "crypto-backend",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJWTSecret())
}
func ParseJWTToken(tokenStr string) (*dto.Claims, error) {
	if tokenStr == "" {
		return nil, jwt.ErrTokenMalformed
	}
	token, err := jwt.ParseWithClaims(tokenStr, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
