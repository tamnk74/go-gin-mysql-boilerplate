package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/tamnk74/todolist-mysql-go/config"
	"github.com/tamnk74/todolist-mysql-go/database"
	itemEntity "github.com/tamnk74/todolist-mysql-go/modules/items"
)

var db *gorm.DB
var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err = database.ConnectDb()
	if err != nil {
		panic("Failed to connect database")
	}
	log.Info("Starting Todolist API server")
	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	itemEntity.RegisterRouters(api, db)
	r.Run(fmt.Sprintf(":%s", config.PORT))
}
