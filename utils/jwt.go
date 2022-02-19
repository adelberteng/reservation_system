package utils

import (
	"fmt"
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var (
	expirationTime = time.Now().Add(5 * time.Minute)
	secretKey = cfg.Section("app").Key("secret_key").String()
)

type Claims struct {
	Payload map[string]string
	jwt.StandardClaims
}

func GenerateJWT(payload map[string]string) (string, error) {
	claims := Claims{
		Payload: map[string]string{"name": "aaa"},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	hmacSecret := []byte(secretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
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
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
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