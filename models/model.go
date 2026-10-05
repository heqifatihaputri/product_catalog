package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock" gorm:"not null"`
}
