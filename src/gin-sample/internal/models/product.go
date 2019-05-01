package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ID    int
	Code  string
	Price uint
}
