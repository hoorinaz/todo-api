package store

import (
	"fmt"
	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/shared/store"
	"github.com/jinzhu/gorm"
	"log"
)


type UserStore struct {
	DB *gorm.DB
}

const(
	tableName="users"
	logger="user-store-GetUser"
)
func(us UserStore) AddUser (u *user.User)error{

db:= us.DB
fmt.Printf("username: %v , password : %v , email: %v ", u.Username , u.Password , u.Email )
 db.Create(&user.User{
	Username: u.Username,
	Email: u.Email,
	Password: u.Password,
	IsRegistered: false,
})
	return nil
}

func (us UserStore) GetUser (u *user.User)error{
	db:= us.DB
	err:= db.Table("users").Where("user_name=?",u.Username).Error
	if err!=nil{
		log.Println(logger,"there is problem to GetUser ",err.Error())
	}
	return nil
}

func NewUserStore () user.UserService{

	//path: shared.store
	s := store.GetDB()
	return UserStore{
		DB: s,
	}

}