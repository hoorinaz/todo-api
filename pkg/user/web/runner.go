package web

import (
	"github.com/gorilla/mux"
	"github.com/hoorinaz/TodoList/pkg/user/processor"
	"github.com/hoorinaz/TodoList/pkg/user/store"
)

func RegisterUserService(r *mux.Router) {
	userStor := store.NewUserStore()
	userProcessor := processor.NewUserProcessor(userStor)
	userWebService := NewUserWebService(userProcessor)

	r.HandleFunc("/user", userWebService.Register).Methods("POST")
}
