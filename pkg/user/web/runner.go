package web

import (
	"github.com/gorilla/mux"
	"github.com/hoorinaz/TodoList/pkg/user/processor"
	"github.com/hoorinaz/TodoList/pkg/user/store"

)

func RegisterUserService (r *mux.Router) {

userStore := store.NewUserStore()
UserProcessor:= processor.NewUserProcessor(userStore)
webService:= NewUserWebService(UserProcessor)


r.HandleFunc("/authentication",webService.Authenticate).Methods("GET")
r.HandleFunc("/register", webService.Register).Methods("POST")

}





