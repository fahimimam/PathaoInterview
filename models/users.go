package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Tags      []Tags `json:"tags"gorm:"many2many:user_tags;"`
}

type Tags struct {
	gorm.Model
	Name   string    `json:"name"`
	Expiry time.Time `json:"expiry"`
}
