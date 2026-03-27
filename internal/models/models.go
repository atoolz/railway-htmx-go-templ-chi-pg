package models

import "time"

type Todo struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}
