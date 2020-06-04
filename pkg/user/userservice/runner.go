package userservice

import (
	"github.com/gorilla/mux"
)

func RegisterUserService(r *mux.Router) {

	userStore := NewUserStore()
	UserProcessor := NewUserProcessor(userStore)
	webService := NewUserWebService(UserProcessor)

	r.HandleFunc("/signin", webService.SignIn).Methods("POST")
	r.HandleFunc("/signup", webService.SignUp).Methods("POST")
}
