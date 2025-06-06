package models

type EventDAO struct {
	ID          int64
	UUID        string
	Title       string
	Description string
	Date        string // RFC3339
}
