package models

import "time"

type ExpenseType struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
