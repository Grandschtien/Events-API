package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"events-api/core/utils"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint) (string, error) {
	secretKey := os.Getenv("SECRETKEY")
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(utils.AccessTokenRefreshTTL).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func GenerateRefreshToken(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("token length must be greater than 0")
	}

	bytes := make([]byte, length)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
