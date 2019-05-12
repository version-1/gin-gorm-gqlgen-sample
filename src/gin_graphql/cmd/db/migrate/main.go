package main

import (
	"gin_graphql/internal/models"
	connection "gin_graphql/pkg/database"
)

func main() {
	db := connection.GetInstance()
	defer connection.Close()

	db.AutoMigrate(&models.User{}, &models.Todo{})
}
