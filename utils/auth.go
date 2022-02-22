package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	expirationTime = time.Now().Add(5 * time.Minute)
	secretKey      = cfg.Section("app").Key("secret_key").String()
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
	if err != nil {
		return false
	}

	return true
}

type Claims struct {
	Payload map[string]string
	jwt.StandardClaims
}

func GenerateJWT(payload map[string]string) (string, error) {
	claims := Claims{
		Payload: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	hmacSecret := []byte(secretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	hmacSecret := []byte(secretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	fmt.Printf("%+v \n", token)
	if !token.Valid {
		fmt.Println("token invalid.")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
