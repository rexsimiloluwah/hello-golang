package utils

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Hashing password using bcrypt
func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, 12)

	if err != nil {
		return "", errors.Wrap(err, "Error hashing password.")
	}

	return string(hashedPassword), nil
}

// Comparing hashed password
func ComparePasswordHash(password string, hash string) error {
	passwordBytes := []byte(password)
	hashedPasswordBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)

	return errors.Wrap(err, "Error comparing password and hash.")
}
