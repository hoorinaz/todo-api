package userservice

import (
	"github.com/gorilla/mux"
)

func RegisterUserService(r *mux.Router) {

	userStore := NewUserStore()
	UserProcessor := NewUserProcessor(userStore)
	webService := NewUserWebService(UserProcessor)

	r.HandleFunc("/authentication", webService.Authenticate).Methods("POST")
	r.HandleFunc("/register", webService.Register).Methods("POST")

}
