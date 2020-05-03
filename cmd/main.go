package main

import (
	"github.com/gorilla/mux"
	userWebService "github.com/hoorinaz/TodoList/pkg/user/web"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	userWebService.RegisterUserService(router)

	log.Panic(http.ListenAndServe(":8080", router))
}
