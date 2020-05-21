package web

import (
	"github.com/gorilla/mux"
	"github.com/hoorinaz/TodoList/pkg/user/processor"
)

func RegisterUserService(r *mux.Router) {
	userProcessor := processor.NewUserProcessor(nil)
	userWebService := NewUserWebService(userProcessor)

	r.HandleFunc("/user", userWebService.Register).Methods("POST")
}
