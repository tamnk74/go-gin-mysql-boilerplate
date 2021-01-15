package models

// Item type
type User struct {
	ID     int `gorm:"primary_key"`
	Name   string
	Gender int
}
