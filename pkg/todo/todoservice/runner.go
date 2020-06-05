package todoservice

import (
	"github.com/gorilla/mux"
	"github.com/hoorinaz/todo-api/shared/gateway"
)

func AddTodoService(r *mux.Router) {
	newMW := gateway.NewMiddleware()

	todoStore := NewTodoStore()
	todoProcessor := NewTodoProcessor(todoStore)
	todoWebServer := NewTodoWebService(todoProcessor)

	r.HandleFunc("/todo/add", newMW.AuthMidd(todoWebServer.AddTodo)).Methods("POST")
	r.HandleFunc("/todo/view", newMW.AuthMidd(todoWebServer.ViewTodo)).Methods("GET")
	r.HandleFunc("/todo/edit", newMW.AuthMidd(todoWebServer.EditTodo)).Methods("PATCH")
	r.HandleFunc("/todo/all", newMW.AuthMidd(todoWebServer.ListTodo)).Methods("GET")
	r.HandleFunc("/todo/delete", newMW.AuthMidd(todoWebServer.DeleteTodo)).Methods("DELETE")

}
