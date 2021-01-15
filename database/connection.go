package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "github.com/tamnk74/todolist-mysql-go/config"
)

// ConnectDb connect to mysql
func ConnectDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
