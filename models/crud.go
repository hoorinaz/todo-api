package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// id
func ViewTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	params := mux.Vars(r)
	Id := params["id"]
	db := Getdb()
	defer db.Close()
	db.First(&todo, Id)
	fmt.Println("FOUND: ", todo)
	json.NewEncoder(w).Encode(todo)
}

//title
func EditTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	params := mux.Vars(r)
	title := params["title"]
	db := Getdb()
	defer db.Close()
	db.First(&todo, 12)
	fmt.Println(todo)
	db.Model(&todo).Update("Title", title)

}

//id
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
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
func AddTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	title := params["title"]
	desc := params["desc"]
	userId := params["userId"]
	intUserId, _ := strconv.Atoi(userId)
	db := Getdb()
	defer db.Close()
	db.Table("todos").Create(&Todo{
		Title:       title,
		Description: desc,
		IsDone:      true,
		UserId:      intUserId,
	})
	fmt.Fprint(w, "Todo add susseccfully. named: ", title)
}

func ViewAll(w http.ResponseWriter, r *http.Request) {
	var allTodo []Todo
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
