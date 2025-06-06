package db

import (
	"database/sql"
	"events-api/models"
	"fmt"
)

type DB struct {
	DB *sql.DB
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func (db *DB) GetEvent(uuid string) (models.Event, error) {
	var event models.EventDAO

	row := db.DB.QueryRow("SELECT * FROM public.events WHERE uuid = $1", uuid)

	if err := row.Scan(&event.ID, &event.UUID, &event.Title, &event.Description, &event.Date); err != nil {
		if err == sql.ErrNoRows {
			return models.Event{}, sql.ErrNoRows
		}

		return models.Event{}, fmt.Errorf("albumsById %s: %v", uuid, err)
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

	rows, err := db.DB.Query("SELECT * FROM public.events")

	if err != nil {
		if err == sql.ErrNoRows {
			return []models.Event{}, nil
		}
		return nil, fmt.Errorf("Internal error: %d", err)
	}

	defer rows.Close()

	for rows.Next() {
		var event models.EventDAO

		if err := rows.Scan(&event.ID, &event.UUID, &event.Title, &event.Description, &event.Date); err != nil {
			return nil, fmt.Errorf("internal error: %d", err)
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

	err := db.DB.QueryRow(
		"INSERT INTO public.events (uuid, title, description, date) VALUES ($1, $2, $3, $4) RETURNING id",
		event.UUID, event.Title, event.Description, event.Date,
	).Scan(&lastId)

	if err != nil {
		return -1, fmt.Errorf("internal error: %d", err)
	}

	return lastId, nil
}

func (db *DB) DeleteEvent(uuid string) error {
	_, err := db.DB.Exec("DELETE FROM public.events WHERE uuid = $1", uuid)

	if err != nil {
		return fmt.Errorf("failed to delete event: %v", err)
	}

	return nil
}
