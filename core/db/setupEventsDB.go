package db

import (
	"database/sql"
	"events-api/core/utils"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func SetupEventDB() (error, *DB) {
	dbname := "eventsdb"
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
	insertEvent := utils.PrepareEvent(
		db,
		"INSERT INTO public.events (uuid, title, description, date) VALUES ($1, $2, $3, $4) RETURNING id",
	)
	getAllEvent := utils.PrepareEvent(
		db,
		"SELECT * FROM public.events",
	)
	getByUUIDEvent := utils.PrepareEvent(
		db,
		"SELECT * FROM public.events WHERE uuid = $1",
	)
	deleteEvent := utils.PrepareEvent(
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
