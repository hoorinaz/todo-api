package main

import (
	"fmt"
	//"github.com/hoorinaz/TodoList/controller/user"
	//"github.com/hoorinaz/TodoList/models"
	"github.com/hoorinaz/TodoList/pkg/user"

	"github.com/hoorinaz/TodoList/pkg/user/web"
	"github.com/hoorinaz/TodoList/shared/store"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db := store.GetDB()
	//db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&user.User{})
	//db.Model(&models.Todo{}).AddForeignKey("user_id", ~"users(id)", "RESTRICT", "RESTRICT")

	router := mux.NewRouter()
	router.HandleFunc("/user",web.RegisterUserService().Register).Methods("POST")
	//router.HandleFunc("/register", user.Register).Methods("POST")
	//router.HandleFunc("/todo/add", auth.Middleware(todo.AddTodo)).Methods("POST")
	//router.HandleFunc("/todo/view/{id}", auth.Middleware(todo.ViewTodo)).Methods("GET")
	//router.HandleFunc("/todo/edit/{id}", auth.Middleware(todo.EditTodo)).Methods("PATCH")
	//router.HandleFunc("/todo/delete/{Id}", auth.Middleware(todo.DeleteTodo)).Methods("DELETE")
	//router.HandleFunc("/todo", auth.Middleware(todo.ViewAll)).Methods("GET")
	//router.HandleFunc("/todo/list", auth.Middleware(todo.GetTodo)).Methods("GET")
	//router.HandleFunc("/authenticate", user.Authenticate).Methods("POST")

	fmt.Println("connect to db")
	http.ListenAndServe(":8080", router)
}
