package store

import (
	"github.com/gocraft/dbr/v2"
	"github.com/hoorinaz/TodoList/pkg/user"
	"log"
)

const (
	tableName  = "account.users"
	loggerName = "account.users-logger "
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
	db     *dbr.Session
	logger log.Logger
}

func (us *UserService) AddUser(u *user.User) error {
	if _, err := us.db.
		NewSession(nil).
		InsertInto(tableName).
		Columns(columns...).
		Record(u).
		Exec(); err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUser(u *user.User) error {
	selectStmt := us.db.NewSession(nil).
		Select(columns...).
		From(tableName)

	if u.Email != "" {
		selectStmt = selectStmt.Where("email = ? ", u.Email)
	} else {
		selectStmt = selectStmt.Where("id = ?", u.ID)
	}

	if err := selectStmt.LoadOne(u); err != nil {
		us.logger.Println(loggerName,
			"error selecting from users table ",
			"error: ",
			err.Error(),
		)

		return err
	}

	return nil
}

func NewUserService(s *dbr.Session) user.UserService {
	return &UserService{
		db:     s,
		logger: log.Logger{},
	}
}
