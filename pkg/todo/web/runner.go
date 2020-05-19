package web

import (
	"github.com/gorilla/mux"
	"github.com/hoorinaz/TodoList/pkg/todo/processor"
	"github.com/hoorinaz/TodoList/pkg/todo/store"
)

func AddTodoService (r *mux.Router){


	todoStore:= store.NewTodoStore()
	todoProcessor:= processor.NewTodoProcessor(todoStore)
	todoWebServer:=NewTodoWebService(todoProcessor)

r.HandleFunc("/todo/add",todoWebServer.AddTodo).Methods("POST")

}
