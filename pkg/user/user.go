package user

import "time"

type User struct {
	ID        int64
	UserName  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserService interface {
	AddUser(*User) error
	GetUser(*User) error
}
