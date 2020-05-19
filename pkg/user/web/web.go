package web

import (
	"encoding/json"
	"fmt"
	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/shared/errorz"
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

	if err=uws.UserProcessor.AddUser(user); err!=nil{
		log.Println(logger,"error in processor layer", err.Error())
		errorz.WriteHttpError(w,http.StatusInternalServerError)
		return
	}

}


func (uwb UserWebService)Authenticate (w http.ResponseWriter, r *http.Request){
	var user *user.User

	err:= json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		errorz.WriteHttpError(w, http.StatusBadRequest,"Bad Request")
		log.Println(logger,"json decode error is: ", err.Error())
		return
	}

	if err :=uwb.UserProcessor.GetUser(user); err!=nil{
		log.Println(logger,"error in processor layer", err.Error())
		errorz.WriteHttpError(w,http.StatusInternalServerError)
		return
	}
fmt.Fprint(w, user)

}

func NewUserWebService (userProcessor user.UserService) UserWebService{

	return UserWebService{
		UserProcessor: userProcessor,
	}
}


