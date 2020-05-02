package user

import "time"

type User struct {
	ID        uint      `db:"id"`
	UserName  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserService interface {
	AddUser(*User) error
	GetUser(*User) error
}
