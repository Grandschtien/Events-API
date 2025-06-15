package db

import (
	"database/sql"
)

type UsersDB struct {
	DB         *sql.DB
	insertStmt *sql.Stmt
	selectStmt *sql.Stmt
}

func (db *UsersDB) SaveUser(tx *sql.Tx, username string, password string) (int, error) {
	var id int

	err := tx.Stmt(db.insertStmt).QueryRow(username, password).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (db *UsersDB) GetUser(userId uint64) (string, error) {
	return "nil", nil
}

func (db *UsersDB) Close() error {
	return db.DB.Close()
}
