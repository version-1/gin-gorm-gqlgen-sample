package main

import (
	"gin_graphql/internal/models"
	"gin_graphql/pkg"
)

func main() {
	db := pkg.Connect("development")
	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&models.User{}, &models.Todo{})
}
