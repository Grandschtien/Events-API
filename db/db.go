package db

import (
	"database/sql"
	"events-api/models"
	"fmt"
)

type DB struct {
	DB             *sql.DB
	insertEvent    *sql.Stmt
	getAllEvent    *sql.Stmt
	getEventWithId *sql.Stmt
	deleteEvent    *sql.Stmt
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func (db *DB) GetEvent(uuid string) (models.Event, error) {
	var event models.EventDAO

	row := db.getEventWithId.QueryRow(uuid)

	if err := row.Scan(&event.ID, &event.UUID, &event.Title, &event.Description, &event.Date); err != nil {
		if err == sql.ErrNoRows {
			return models.Event{}, sql.ErrNoRows
		}

		return models.Event{}, fmt.Errorf("albumsById %s: %w", uuid, err)
	}

	eventDTO := models.Event{
		UUID:        event.UUID,
		Title:       event.Title,
		Description: event.Description,
		Date:        event.Date,
	}

	return eventDTO, nil
}

func (db *DB) GetAllEvents() ([]models.Event, error) {
	events := make([]models.Event, 0)

	rows, err := db.getAllEvent.Query()

	if err != nil {
		if err == sql.ErrNoRows {
			return []models.Event{}, nil
		}
		return nil, fmt.Errorf("Internal error: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var event models.EventDAO

		if err := rows.Scan(&event.ID, &event.UUID, &event.Title, &event.Description, &event.Date); err != nil {
			return nil, fmt.Errorf("internal error: %w", err)
		}

		eventDTO := models.Event{
			UUID:        event.UUID,
			Title:       event.Title,
			Description: event.Description,
			Date:        event.Date,
		}

		events = append(events, eventDTO)
	}

	return events, nil
}

func (db *DB) SaveEvent(event models.Event) (int64, error) {
	var lastId int64

	err := db.insertEvent.QueryRow(event.UUID, event.Title, event.Description, event.Date).Scan(&lastId)

	if err != nil {
		return -1, fmt.Errorf("internal error: %w", err)
	}

	return lastId, nil
}

func (db *DB) DeleteEvent(uuid string) error {
	_, err := db.deleteEvent.Exec(uuid)

	if err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}

	return nil
}
