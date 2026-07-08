package models

import "time"

type Customer struct {
	ID int
	FirstName string
	LastName string
	Email string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
