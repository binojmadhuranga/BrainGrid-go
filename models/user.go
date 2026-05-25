package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"`

	Role string `json:"role" gorm:"default:USER"`

	Status string `json:"status" gorm:"default:ACTIVE"`
}