package web

import (
	"github.com/gorilla/mux"
	"github.com/hoorinaz/TodoList/pkg/todo/processor"
	"github.com/hoorinaz/TodoList/pkg/todo/store"
	"github.com/hoorinaz/TodoList/shared/gateway"
)

func AddTodoService(r *mux.Router) {
	newMW := gateway.NewMiddleware()

	todoStore := store.NewTodoStore()
	todoProcessor := processor.NewTodoProcessor(todoStore)
	todoWebServer := NewTodoWebService(todoProcessor)

	r.HandleFunc("/todo/add", newMW.Authenticate(todoWebServer.AddTodo)).Methods("POST")
	r.HandleFunc("/todo/view", newMW.Authenticate(todoWebServer.ViewTodo)).Methods("GET")
	r.HandleFunc("/todo/edit", newMW.Authenticate(todoWebServer.EditTodo)).Methods("PATCH")
	r.HandleFunc("/todo/all", newMW.Authenticate(todoWebServer.ListTodo)).Methods("GET")
	r.HandleFunc("/todo/delete", newMW.Authenticate(todoWebServer.DeleteTodo)).Methods("DELETE")

}
