package user

import (
	"context"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsRegistered bool
	Token        string
}

type UserService interface {
	AddUser(context.Context, *User) error
	GetUser(context.Context, *User) error
}
