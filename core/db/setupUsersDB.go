package db

import (
	"database/sql"
	"events-api/core/utils"
	"fmt"
	"os"
)

func SetupUserDB() (error, *UsersDB) {
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
		"INSERT INTO public.users (username, password) VALUES ($1, $2) RETURNING id",
	)
	selectStmt := utils.PrepareEvent(
		db,
		"SELECT * FROM public.users WHERE username = $1",
	)
	fmt.Println("Successfully connected!")

	return nil, &UsersDB{
		DB:         db,
		insertStmt: insertStmt,
		selectStmt: selectStmt,
	}
}
