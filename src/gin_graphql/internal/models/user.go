package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        int
	Name      string
	Todos     []Todo
	CreatedAt time.Time
	UpdatedAt time.Time
}
