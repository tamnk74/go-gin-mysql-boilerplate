package models

// Item type
type Item struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"unique,not null"`
	description string
	Status      int
}
