package models

type EventDAO struct {
	UUID        string
	Title       string
	Description string
	Date        string // RFC3339
}
