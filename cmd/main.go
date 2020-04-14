package main

import (
	"fmt"
	"github.com/hoorinaz/TodoList/todo"
	"github.com/hoorinaz/TodoList/user"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db := todo.Getdb()
	db.AutoMigrate(&todo.Todo{})
	db.AutoMigrate(&user.User{})
	//db.Model(&todo.Todo{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")

	router := mux.NewRouter()
	router.HandleFunc("/", todo.ViewAll).Methods("GET")
	router.HandleFunc("/todo/view/{id}", todo.ViewTodo).Methods("GET")
	router.HandleFunc("/todo/edit/{titlaccountse}", todo.EditTodo).Methods("PUT")
	router.HandleFunc("/todo/delete/{Id}", todo.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todo/add/{title}/{desc}/{userId}", todo.AddTodo).Methods("POST")
	router.HandleFunc("/user/add/{username}/{email}", user.AddUser).Methods("POST")
	router.HandleFunc("/user/todos", user.GetTodo).Methods("GET")
	router.HandleFunc("/signup", user.Signup).Methods("POST")
	router.HandleFunc("/signin", user.Signin).Methods("POST")

	fmt.Println("connect to db")
	http.ListenAndServe(":8080", router)
}
