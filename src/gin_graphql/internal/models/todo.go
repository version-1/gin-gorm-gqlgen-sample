package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Todo struct {
	ID        int
	Text      string
	Done      bool
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GTodo struct {
	gorm.Model
	Todo
}
