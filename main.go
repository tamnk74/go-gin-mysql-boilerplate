package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/tamnk74/todolist-mysql-go/config"
	"github.com/tamnk74/todolist-mysql-go/database"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	"github.com/tamnk74/todolist-mysql-go/router"
	"github.com/tamnk74/todolist-mysql-go/schedulers"
	"github.com/tamnk74/todolist-mysql-go/utils/redis"
)

var db *gorm.DB
var err error

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.Connect()
	if err != nil {
		panic("Failed to connect database")
	}

	redis.Init()

	log.Info("Starting API server at port " + config.PORT)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middlewares.HandleApiError())
	router.Init(r)
	schedulers.Init()
	r.Run(":" + config.PORT)
}
