package utils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(tokenString string) (jwt.Claims, error) {
	secretKey := os.Getenv("SECRETKEY")
	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
