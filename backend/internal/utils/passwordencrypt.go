package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt.
// It returns the hashed password as a string and any error encountered.
// The cost parameter determines the computational cost of the hashing process.
func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, 12)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
