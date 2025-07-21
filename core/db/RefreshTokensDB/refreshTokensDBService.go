package db

import (
	"database/sql"
	"events-api/authentication/models"
	"events-api/core/utils"
	"fmt"
	"time"
)

func (db *RefreshTokensDB) GetRefreshTokens(userID int) (error, []models.RefreshTokenDAO) {
	tokens := make([]models.RefreshTokenDAO, 0)

	rows, err := db.selectStmt.Query(userID, false)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, []models.RefreshTokenDAO{}
		}
		return fmt.Errorf("Internal error: %w", err), nil
	}

	defer rows.Close()

	for rows.Next() {
		var token models.RefreshTokenDAO
		if err := rows.Scan(
			&token.ID,
			&token.UserID,
			&token.TokenHash,
			&token.IssuedAt,
			&token.ExpiresAt,
			&token.Revoked,
		); err == nil {
			tokens = append(tokens, token)
		}
	}
	return nil, tokens
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

func (db *RefreshTokensDB) DeleteRefreshToken(userID int64) error {
	return nil
}

func (db *RefreshTokensDB) RevokeRefreshToken(userID int, revoked bool) error {
	_, err := db.revokeTokenStmt.Query(revoked, userID)

	if err != nil {
		return err
	}

	return nil
}
