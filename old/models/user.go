package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsRegistered bool
	Token        string
}
