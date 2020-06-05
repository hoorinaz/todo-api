package main

import (
	"fmt"

	"github.com/hoorinaz/todo-api/pkg/todo"
	"github.com/hoorinaz/todo-api/pkg/todo/todoservice"
	"github.com/hoorinaz/todo-api/pkg/user/userservice"

	//"github.com/hoorinaz/TodoList/controller/user"
	//"github.com/hoorinaz/TodoList/models"
	"github.com/hoorinaz/todo-api/pkg/user"

	"net/http"

	"github.com/hoorinaz/todo-api/shared/connection"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db := connection.GetDB()
	db.AutoMigrate(&todo.Todo{})
	db.AutoMigrate(&user.User{})
	db.Model(&todo.Todo{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	router := mux.NewRouter()

	userservice.RegisterUserService(router)
	todoservice.AddTodoService(router)

	fmt.Println("connect to db")

	http.ListenAndServe(":8080", router)
}
