package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) (string, error) {
	pw := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(pw, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func VerifyPassword(password, passwordHash string) bool {
	pw := []byte(password)
	hash := []byte(passwordHash)

	err := bcrypt.CompareHashAndPassword(hash, pw)

	return err == nil
}
