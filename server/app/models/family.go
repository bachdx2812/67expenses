package models

import "time"

type Family struct {
	ID        int32
	Name      string
	Users     []User
	CreatedAt time.Time
	UpdatedAt time.Time
}
