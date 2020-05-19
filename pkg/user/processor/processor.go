package processor

import (
	"github.com/hoorinaz/TodoList/pkg/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const logger ="user-process"

type UserProcessor struct {

	UserStore user.UserService
}

func (up UserProcessor) AddUser(u *user.User) error{

	if len(u.Password ) < 6 {
		//fmt.Fprint(w,"password is less than 6 character")
		log.Println(logger, "password is less than 6 character")
		return nil
	}
	hashPass, err := HashPassword(u.Password)
	if err!=nil{
		log.Println(logger, "Error in Password Hashing: ", err.Error())
		//errorz.WriteHttpError(w, http.StatusBadRequest)
		return nil

	}
	u.Password=hashPass
	return up.UserStore.AddUser(u)
}


func(up UserProcessor) GetUser (u *user.User)error{
	return up.UserStore.GetUser(u)
}

func NewUserProcessor (userStore user.UserService) user.UserService{
	return UserProcessor{
		UserStore: userStore,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}