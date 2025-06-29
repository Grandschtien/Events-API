package db

import (
	"database/sql"
	"events-api/authentication/models"
	"fmt"
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

func (db *UsersDB) GetUser(username string) (models.LoginUserDAO, error) {
	var user models.LoginUserDAO

	row := db.selectStmt.QueryRow(username)

	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CrearedAt); err != nil {
		return models.LoginUserDAO{}, fmt.Errorf("internal error: %w", err)
	}

	return user, nil
}

func (db *UsersDB) Close() error {
	return db.DB.Close()
}
