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
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}
