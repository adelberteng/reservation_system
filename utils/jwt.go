package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	expirationTime = time.Now().Add(30 * time.Minute)
	secretKey      = Conf.App.SecretKey
)

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

func ParseJWT(tokenString string) (*jwt.Token, error) {
	hmacSecret := []byte(secretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func VerifyJWT(token *jwt.Token) bool {
	return token.Valid
}

func RetrieveJWT(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("token retrieve not complete.")
	}
}
