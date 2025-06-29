package models

import "time"

type LoginUserDAO struct {
	ID        int
	Username  string
	Password  []byte
	CrearedAt time.Time
}

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
