package models

import "time"

type Ticket struct {
	ID          int
	Title       string `binding:"required"`
	Category    string `binding:"required"`
	Description string `binding:"required"`
	Status      string
	Client      int
	Assign      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
