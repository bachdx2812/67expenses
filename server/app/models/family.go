package models

import "time"

type Family struct {
	Id        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
