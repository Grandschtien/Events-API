package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
)

func SetupEventDB() (error, *DB) {
	password := os.Getenv("PSQLDBPASS")
	user := os.Getenv("PSQLDBUSER")
	dbname := "eventsdb"

	info := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", info)

	if err != nil {
		return err, nil
	}

	err = db.Ping()
	if err != nil {
		return err, nil
	}

	insertEvent := prepareEvent(
		db,
		"INSERT INTO public.events (uuid, title, description, date) VALUES ($1, $2, $3, $4) RETURNING id",
	)
	getAllEvent := prepareEvent(
		db,
		"SELECT * FROM public.events",
	)
	getByUUIDEvent := prepareEvent(
		db,
		"SELECT * FROM public.events WHERE uuid = $1",
	)
	deleteEvent := prepareEvent(
		db,
		"DELETE FROM public.events WHERE uuid = $1",
	)

	fmt.Println("Successfully connected!")

	return nil, &DB{
		DB:             db,
		insertEvent:    insertEvent,
		getAllEvent:    getAllEvent,
		getEventWithId: getByUUIDEvent,
		deleteEvent:    deleteEvent,
	}
}

func prepareEvent(
	conn *sql.DB,
	operation string,
) *sql.Stmt {
	stmt, err := conn.Prepare(
		operation,
	)

	if err != nil {
		panic("Operation did not prepare")
	}

	return stmt
}
