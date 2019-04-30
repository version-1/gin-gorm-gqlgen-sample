package main

import (
	"gin-sample/internal/models"
	"gin-sample/pkg"
)

func main() {
	db := pkg.Connect("development")
	defer db.Close()

	db.AutoMigrate(&models.Product{})
	db.Create(&models.Product{Code: "L1212", Price: 1000})
	db.Create(&models.Product{Code: "ABABA", Price: 2000})
	db.Create(&models.Product{Code: "CDCDC", Price: 3000})
}
