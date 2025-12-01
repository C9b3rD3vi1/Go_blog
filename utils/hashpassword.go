package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the user's password
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    password = string(hashedPassword)
    return password, nil
}