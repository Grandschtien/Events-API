package db

import (
	"database/sql"
	"events-api/models"
)

type DB struct {
	DB *sql.DB
}

func (db *DB) GetEvent(uuid string) (models.Event, error) {

}

func (db *DB) GetAllEvents() ([]models.Event, error) {

}

func (db *DB) SaveEvent(event models.EventDAO) (int64, error) {

}

func (db *DB) DeleteEvent(uuid string) error {

}

func (db *DB) AddEvent() error {

}
