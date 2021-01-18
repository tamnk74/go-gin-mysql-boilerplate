package main

import (
	"github.com/tamnk74/todolist-mysql-go/database"
	"github.com/tamnk74/todolist-mysql-go/models"
)

func main() {
	err := database.Connect()
	if err != nil {
		panic("Failed to connect database")
	}
	db := database.GetDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Item{})
}
