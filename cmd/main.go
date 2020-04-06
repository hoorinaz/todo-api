package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hoorinaz/TodoList/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	// db := models.Getdb()
	// db.Model(&models.Todo{}).AddForeignKey("user_id", "accounts(id)", "RESTRICT", "RESTRICT")
	// db.Model(&models.Account{}).Related(&models.Todo{})

	router := mux.NewRouter()
	router.HandleFunc("/", models.ViewAll).Methods("GET")
	router.HandleFunc("/view/{id}", models.ViewTodo).Methods("GET")
	router.HandleFunc("/edit/{title}", models.EditTodo).Methods("PUT")
	router.HandleFunc("/delete/{Id}", models.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/add/{title}/{desc}/{userId}", models.AddTodo).Methods("POST")
	router.HandleFunc("/adduser/{username}/{email}", models.AddUser).Methods("POST")
	router.HandleFunc("/signin", models.Signin).Methods("POST")
	router.HandleFunc("/login", models.Login).Methods("GET")

	fmt.Println("connect to db")
	http.ListenAndServe(":8080", router)
}
