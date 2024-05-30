package models

import "time"

type User struct {
	ID                uint
	Phone             string
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
