package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	Text   string
	Done   bool
	UserID int
}

type ExtendedTodo struct {
	gorm.Model
	Todo
}
