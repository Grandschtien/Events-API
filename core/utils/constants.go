package utils

import "time"

const (
	Host = "localhost"
	Port = 5432
)

const AccessTokenRefreshTTL = 24 * time.Hour
const RefreshTokenRefreshTTL = 30 * 24 * time.Hour
