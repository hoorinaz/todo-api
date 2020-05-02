package store

import (
	"github.com/gocraft/dbr/v2"
	"github.com/hoorinaz/TodoList/pkg/user"
)

const (
	tableName = "account.users"
)

var columns = []string{
	"id",
	"username",
	"email",
	"password",
	"created_at",
	"updated_at",
}

type UserService struct {
	db *dbr.Session
}

func (us *UserService) AddUser(u *user.User) error {
	if _, err := us.db.
		NewSession(nil).
		InsertInto(tableName).
		Columns(columns...).Exec(); err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUser(u *user.User) error {
	return nil
}

func NewUserService(s *dbr.Session) user.UserService {
	return &UserService{
		db: s,
	}
}
