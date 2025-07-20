package db

import (
	"events-api/authentication/models"
	"events-api/core/utils"
	"time"
)

func (db *RefreshTokensDB) GetRefreshToken(userID int) (error, models.RefreshTokenDAO) {
	var token models.RefreshTokenDAO

	row := db.selectStmt.QueryRow(userID)

	if err := row.Scan(
		&token.ID,
		&token.UserID,
		&token.TokenHash,
		&token.IssuedAt,
		&token.ExpiresAt,
		&token.Revoked,
	); err != nil {
		return err, models.RefreshTokenDAO{}
	}
	return nil, token
}

func (db *RefreshTokensDB) SaveRefreshToken(userID int, token string) error {
	var id int

	now := time.Now()
	expiresAt := now.Add(utils.RefreshTokenRefreshTTL)

	err := db.insertStmt.QueryRow(userID, token, now, expiresAt, false).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}
