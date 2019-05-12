package models

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	Todos     []Todo
	CreatedAt time.Time
	UpdatedAt time.Time
}
