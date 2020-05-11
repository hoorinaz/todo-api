package web

import (
	"encoding/json"
	"fmt"
	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/shared/errorz"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

const logger ="user-web"
type UserWebService struct {
	UserProcessor 	user.UserService
}

func (uws UserWebService) Register(w http.ResponseWriter, r *http.Request){

	var user *user.User
	err:=json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		errorz.WriteHttpError(w, http.StatusBadRequest,"Bad Request")
		log.Println(logger,"json decode error is: ", err.Error())
		return
	}
	if len(user.Password ) < 6 {
		fmt.Fprint(w,"password is less than 6 character")
		return
	}
	hashPass, err := HashPassword(user.Password)
	if err!=nil{
		log.Println(logger,"Error in Password Hashing: ", err.Error())
		errorz.WriteHttpError(w, http.StatusBadRequest)
		return

	}
	user.Password=hashPass

	if err=uws.UserProcessor.AddUser(user); err!=nil{
		log.Println(logger,"error in processor layer", err.Error())
		errorz.WriteHttpError(w,http.StatusInternalServerError)
		return
	}

}


func Authenticate (w http.ResponseWriter, r *http.Request){
	var user user.User

	err:= json.NewDecoder(r.Body).Decode(user)
	if err!=nil{
		errorz.WriteHttpError(w, http.StatusBadRequest,"Bad Request")
		log.Println(logger,"json decode error is: ", err.Error())
		return
	}
	fmt.Fprint(w, user)
}

func NewUserWebService (userProcessor user.UserService) UserWebService{

	return UserWebService{
		UserProcessor: userProcessor,
	}
}


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}