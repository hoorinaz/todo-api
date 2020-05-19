package store

import (
	"fmt"
	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/shared/connection"
	"github.com/jinzhu/gorm"
	"log"
)


type UserStore struct {
	DB *gorm.DB
}

const(
	tableName="users"
	logger="user-connection-GetUser"
)
func(us UserStore) AddUser (u *user.User)error{

db:= us.DB
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
	var dbUser user.User
	err:= db.Table("users").Where("username=?",u.Username).First(&dbUser).Error
	if err!=nil{
		log.Println(logger,"there is problem to GetUser ",err.Error())
	}
	fmt.Printf("username: %v , password : %v , email: %v ", dbUser.Username , dbUser.Password , dbUser.Email )

	return nil
}

func NewUserStore () user.UserService{

	//path: shared.connection
	s := connection.GetDB()
	return UserStore{
		DB: s,
	}

}