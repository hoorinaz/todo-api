package userservice

import (
	"context"
	"log"

	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/shared/connection"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	DB *gorm.DB
}

type UserStoreInterface interface {
	AddUser(context.Context, *user.User) error
	GetUser(context.Context, *user.User) error
}

const (
	tableName   = "users"
	loggerStore = "user-connection-GetUser"
)

func (us UserStore) AddUser(ctx context.Context, u *user.User) error {

	db := us.DB
	db.Create(&user.User{
		Username:     u.Username,
		Email:        u.Email,
		Password:     u.Password,
		IsRegistered: false,
	})
	return nil
}

func (us UserStore) GetUser(ctx context.Context, u *user.User) error {
	db := us.DB

	err := db.Table("users").Where("username=?", u.Username).First(&u).Error
	if err != nil {
		log.Println(logger, "there is problem to GetUser ", err.Error())
		return err
	}
	// fmt.Printf("username: %v , password : %v , email: %v ", u.Username, u.Password, u.Email)

	return nil
}

func NewUserStore() UserStoreInterface {

	//path: shared.connection
	s := connection.GetDB()
	return UserStore{
		DB: s,
	}

}
