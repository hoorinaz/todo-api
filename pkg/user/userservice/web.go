package userservice

import (
	"context"
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
	Register(context.Context, *user.User) error
	Authenticate(context.Context, *user.User) error
}

func (uws UserWebService) Register(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	u := new(user.User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		log.Println(logger, "json decode error is: ", err.Error())
		return
	}

	if err = uws.UserProcessor.SignUp(ctx, u); err != nil {
		log.Println(logger, "error in processor layer", err.Error())
		errorz.WriteHttpError(w, http.StatusInternalServerError)
		return
	}

}

func (uws UserWebService) Authenticate(w http.ResponseWriter, r *http.Request) {
	u := new(user.User)

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		log.Println(logger, "json decode error is: ", err.Error())
		return
	}

	ctx := r.Context()
	err = uws.UserProcessor.SignIn(ctx, u)
	if err != nil {
		log.Println(logger, "error in processor layer", err)
		errorz.WriteHttpError(w, http.StatusInternalServerError)
		return
	}

}

func NewUserWebService(userProcessor UserProcessorInterface) UserProcessorInterface {

	return UserWebService{
		UserProcessor: userProcessor,
	}
}
