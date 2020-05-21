package store

import (
	"context"
	"log"
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/shared/store"
	_ "github.com/lib/pq"
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

type (
	prepareData func(*user.User)
	UserStore   struct {
		DB             *dbr.Session
		PrepareAddData prepareData
	}
)

func (us *UserStore) AddUser(ctx context.Context, u *user.User) error {
	us.PrepareAddData(u)
	if _, err := us.DB.
		NewSession(nil).
		InsertInto(tableName).
		Columns(columns[1:]...).
		Record(u).
		Exec(); err != nil {
		return err
	}

	return nil
}

func (us *UserStore) GetUser(ctx context.Context, u *user.User) error {
	selectStmt := us.DB.NewSession(nil).
		Select(columns...).
		From(tableName)

	if u.Email != "" {
		selectStmt = selectStmt.Where("email = ? ", u.Email)
	} else {
		selectStmt = selectStmt.Where("id = ?", u.ID)
	}

	if err := selectStmt.LoadOne(u); err != nil {
		log.Println(loggerName,
			"error selecting from users table ",
			"error: ",
			err.Error(),
		)

		return err
	}
	return nil
}

func (us *UserStore) Close() error {
	return us.DB.Close()
}

func NewUserStore() user.UserService {
	s := store.NewDB()

	return &UserStore{
		DB: s,
		PrepareAddData: func(u *user.User) {
			now := time.Now()
			u.CreatedAt = now
			u.UpdatedAt = now
		},
	}
}
