package utils

import "golang.org/x/crypto/bcrypt"

func CheckPassword(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}
