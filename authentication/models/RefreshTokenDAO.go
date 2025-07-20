package models

import "time"

type RefreshTokenDAO struct {
	ID        int
	UserID    int
	TokenHash string
	IssuedAt  time.Time
	ExpiresAt time.Time
	Revoked   bool
}
