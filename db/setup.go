package db

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	host = "localhost"
	port = 5432
	user = "root"
)

func Setup(dbname string) *sql.DB {
	password := os.Getenv("DBPASS")
	info := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", info)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
