package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Name string
}

type ExtendedUser struct {
	gorm.Model
	User
}
