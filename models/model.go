package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"unique"`
	Password string
	Todos []Todo
}

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	Completed bool `json:"completed"`
	UserID uint
}