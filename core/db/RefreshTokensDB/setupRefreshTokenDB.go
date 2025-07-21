package db

import (
	"database/sql"
	"events-api/core/utils"
	"fmt"
	"os"
)

type RefreshTokensDB struct {
	DB              *sql.DB
	insertStmt      *sql.Stmt
	selectStmt      *sql.Stmt
	revokeTokenStmt *sql.Stmt
}

func (db *RefreshTokensDB) Close() error {
	return db.DB.Close()
}

func SetupRefreshTokensDB() (error, *RefreshTokensDB) {
	dbname := "users"
	var DBName = os.Getenv("PSQLDBUSER")
	var DBPassword = os.Getenv("PSQLDBPASS")

	info := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		utils.Host, utils.Port, DBName, DBPassword, dbname)

	db, err := sql.Open("postgres", info)

	if err != nil {
		return err, nil
	}

	err = db.Ping()

	if err != nil {
		return err, nil
	}

	insertStmt := utils.PrepareEvent(
		db,
		"INSERT INTO public.refresh_tokens (user_id, token_hash, issued_at, expires_at, revoked) VALUES ($1, $2, $3, $4, $5) RETURNING id",
	)
	selectStmt := utils.PrepareEvent(
		db,
		"SELECT * FROM public.refresh_tokens WHERE user_id = $1 AND revoked = $2",
	)

	revokeStmt := utils.PrepareEvent(
		db,
		"UPDATE public.refresh_tokens SET revoked = $1 WHERE user_id = $2",
	)

	fmt.Println("Successfully connected!")

	return nil, &RefreshTokensDB{
		DB:              db,
		insertStmt:      insertStmt,
		selectStmt:      selectStmt,
		revokeTokenStmt: revokeStmt,
	}
}
