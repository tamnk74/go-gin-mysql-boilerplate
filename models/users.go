package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User type
type User struct {
	ID        int     `gorm:"primary_key"`
	Name      string  `gorm:"unique,not null"`
	Email     *string `gorm:"unique,index,not null"`
	Password  string
	Status    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
	return
}
