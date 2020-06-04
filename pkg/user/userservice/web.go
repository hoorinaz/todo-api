package userservice

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/shared/errorz"
)

// const logger = "user-web"

type UserWebService struct {
	UserProcessor UserProcessorInterface
}

type UserWebInterface interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

func (uws UserWebService) SignUp(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	u := new(user.User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		log.Println(logger, "json decode error is: ", err.Error())
		return
	}

	if err = uws.UserProcessor.Register(ctx, u); err != nil {
		log.Println(logger, "error in processor layer", err.Error())
		errorz.WriteHttpError(w, http.StatusInternalServerError)
		return
	}

}

func (uws UserWebService) SignIn(w http.ResponseWriter, r *http.Request) {
	u := new(user.User)

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		log.Println(logger, "json decode error is: ", err.Error())
		return
	}

	ctx := r.Context()
	err = uws.UserProcessor.Authenticate(ctx, u)
	if err != nil {
		log.Println(logger, "error in processor layer", err)
		errorz.WriteHttpError(w, http.StatusInternalServerError)
		return
	}

}

func NewUserWebService(userProcessor UserProcessorInterface) UserWebInterface {

	return UserWebService{
		UserProcessor: userProcessor,
	}
}
