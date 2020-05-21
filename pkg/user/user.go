package user

import (
	"context"
	"time"
)

type User struct {
	ID        int64
	UserName  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserService interface {
	AddUser(context.Context, *User) error
	GetUser(context.Context, *User) error
}
