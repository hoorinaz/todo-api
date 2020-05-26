package main

import (
	"fmt"

	"github.com/hoorinaz/TodoList/pkg/todo"
	todoWeb "github.com/hoorinaz/TodoList/pkg/todo/web"
	userWeb "github.com/hoorinaz/TodoList/pkg/user/web"

	//"github.com/hoorinaz/TodoList/controller/user"
	//"github.com/hoorinaz/TodoList/models"
	"github.com/hoorinaz/TodoList/pkg/user"

	"net/http"

	"github.com/hoorinaz/TodoList/shared/connection"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db := connection.GetDB()
	db.AutoMigrate(&todo.Todo{})
	db.AutoMigrate(&user.User{})
	db.Model(&todo.Todo{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	router := mux.NewRouter()
	userWeb.RegisterUserService(router)
	todoWeb.AddTodoService(router)

	//router.HandleFunc("/todo/add", auth.Middleware(todo.AddTodo)).Methods("POST")
	//router.HandleFunc("/todo/view/{id}", auth.Middleware(todo.ViewTodo)).Methods("GET")
	//router.HandleFunc("/todo/edit/{id}", auth.Middleware(todo.EditTodo)).Methods("PATCH")
	//router.HandleFunc("/todo/delete/{Id}", auth.Middleware(todo.DeleteTodo)).Methods("DELETE")
	//router.HandleFunc("/todo", auth.Middleware(todo.ViewAll)).Methods("GET")
	//router.HandleFunc("/todo/list", auth.Middleware(todo.GetTodo)).Methods("GET")
	//router.HandleFunc("/authenticate", user.Authenticate).Methods("POST")
	//router.HandleFunc("/register", user.Register).Methods("POST")

	fmt.Println("connect to db")

	http.ListenAndServe(":8080", router)
}
