package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	value, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(value), nil
}

func ValidatePassword(hashPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
}
