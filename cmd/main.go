package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hoorinaz/TodoList/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", viewAll).Methods("GET")
	router.HandleFunc("/view/{id}", viewTodo).Methods("GET")
	router.HandleFunc("/edit/{title}", editTodo).Methods("PUT")
	router.HandleFunc("/delete/{Id}", deleteTodo).Methods("DELETE")
	router.HandleFunc("/add/{title}/{desc}/{userId}", addTodo).Methods("POST")
	fmt.Println("connect to db")
	http.ListenAndServe(":8080", router)
}

// id
func viewTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	params := mux.Vars(r)
	Id := params["id"]
	db := Getdb()
	defer db.Close()
	db.First(&todo, Id)
	fmt.Println("FOUND: ", todo)
	json.NewEncoder(w).Encode(todo)
}

//title
func editTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	params := mux.Vars(r)
	title := params["title"]
	db := Getdb()
	defer db.Close()
	db.First(&todo, 12)
	fmt.Println(todo)
	db.Model(&todo).Update("Title", title)

}

//id
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	// db.Delete(&todo)fi
	params := mux.Vars(r)
	Id := params["Id"]
	db := Getdb()
	defer db.Close()
	db.Where("Id =?", Id)

	db.Delete(&todo)
	fmt.Fprint(w, todo)
}

// title/desc/userid
func addTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	title := params["title"]
	desc := params["desc"]
	userId := params["userId"]
	x, _ := strconv.Atoi(userId)

	db := Getdb()
	defer db.Close()
	db.Create(&models.Todo{
		Title:       title,
		Description: desc,
		IsDone:      true,
		UserId:      x,
	})

}

func viewAll(w http.ResponseWriter, r *http.Request) {
	var allTodo []models.Todo
	fmt.Println("dsfsdf")
	db := Getdb()
	defer db.Close()
	db.Find(&allTodo)
	fmt.Fprint(w, allTodo)
}

func Getdb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=hoorzy dbname=hoorzy password=sahar67 sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
