package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
)

func Setup(dbname string) *sql.DB {
	info := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", info)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
