package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	value, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(value), nil
}

func ValidatePassword(hashPassword string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), password)
	return err == nil
}
