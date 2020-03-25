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

// func allTodo (w http.ResponseWriter ,r *http.Request){
// fmt.Println("all todo lists")
// json.NewEncoder(w).Encode(todo)
// }

// type HttpRouter struct{
// 	*mux.Router
// 	*gorm.DB
// }

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

func viewTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	params := mux.Vars(r)
	Id := params["id"]
	fmt.Fprint(w, Id)
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=hoorzy dbname=hoorzy password=sahar67 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db = db.First(&todo)

	fmt.Println("FOUND: ", todo)
	json.NewEncoder(w).Encode(todo)
}

func editTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	params := mux.Vars(r)
	title := params["title"]
	fmt.Fprint(w, title)
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=hoorzy dbname=hoorzy password=sahar67 sslmode=disable")
	if err != nil {
		panic(err)
	}
	// db.Where("title= ?", title)
	db.First(&todo, 4)
	db.Model(&todo).Update("Title", "Hoorie")
	// fmt.Fprintf(w, "Edit Hoorie Nazati")

}
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	// db.Delete(&todo)fi
	params := mux.Vars(r)
	Id := params["Id"]
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=hoorzy dbname=hoorzy password=sahar67 sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.Where("Id =?", Id)

	db.Delete(&todo)
	fmt.Fprint(w, "Delete Todo list items")
}
func addTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	title := params["title"]
	desc := params["desc"]
	userId := params["userId"]
	x, _ := strconv.Atoi(userId)
	// fmt.Println(reflect.TypeOf(x))

	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=hoorzy dbname=hoorzy password=sahar67 sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }
	db := Getdb()
	db.Create(&models.Todo{
		Title:       title,
		Description: desc,
		IsDone:      true,
		UserId:      x,
	})

}
func Getdb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=hoorzy dbname=hoorzy password=sahar67 sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func viewAll(w http.ResponseWriter, r *http.Request) {
	var allTodo []models.Todo
	fmt.Println("dsfsdf")
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=hoorzy dbname=hoorzy password=sahar67 sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.Find(&allTodo)
	fmt.Fprint(w, allTodo)
}
