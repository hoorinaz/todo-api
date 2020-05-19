package todo

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hoorinaz/TodoList/old/models"
	"github.com/hoorinaz/TodoList/shared/auth"
	"github.com/hoorinaz/TodoList/shared/connection"
	"github.com/hoorinaz/TodoList/shared/errorz"
	"log"
	"net/http"
)

// id~
func ViewTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	params := mux.Vars(r)
	id := params["id"]
	db := connection.GetDB()
	user:=auth.GetUserformRequest(w)

	if err:= db.Table("todos").Where("Id =?",id).First(&todo).Error; err!=nil{
		log.Printf("connection error is: %v", err.Error())
		errorz.WriteHttpError(w, http.StatusInternalServerError,  "connection error")
		return
	}
	if user.ID!=todo.UserID{
		errorz.WriteHttpError(w, http.StatusUnauthorized, "you dont have enough permission")
		return
	}

	json.NewEncoder(w).Encode(todo)
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}

//id => param ,title, description  => json
func EditTodo(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserformRequest(w)
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("bad request error is: ", err)
		errorz.WriteHttpError(w, http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	dbTodo := models.Todo{}
	db := connection.GetDB()
	err = db.Table("todos").Where("Id =?", id).Find(&dbTodo).Error
	if err != nil {
		log.Printf("problme found to get todo with %v. error: %v", id, err.Error())
		msg := "todo with" + id + "not found"
		errorz.WriteHttpError(w, http.StatusUnauthorized, msg)
		return
	}
	if dbTodo.ID == 0 {
		fmt.Fprintf(w, "incorrect ID, todo not found")
		return
	}

	if dbTodo.UserID != user.ID {
		errorz.WriteHttpError(w, http.StatusUnauthorized, "you dont have enough permission")
		return
	}

	if res := db.Model(&todo).
		Where("id = ?", dbTodo.ID).Updates(models.Todo{Title: todo.Title,
		Description: todo.Description}); res.Error != nil {
		log.Println("error", res.Error.Error())
	}

	fmt.Fprintf(w, "Todo %v has been changed successfully.", id)
	err = db.Close()
	if err != nil {
		log.Println("error in DB session closing ", err.Error())
	}
}

//id
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserformRequest(w)
	var todo models.Todo
	params := mux.Vars(r)
	Id := params["Id"]
	db := connection.GetDB()
	if res := db.Table("todos").Where("Id =?", Id).Find(&todo); res.Error != nil {
		fmt.Fprintf(w, "connection problem")
		errorz.WriteHttpError(w, http.StatusInternalServerError)
		return
	}
	if todo.UserID != user.ID {
		errorz.WriteHttpError(w, http.StatusUnauthorized, "you are not authorized")
		return
	}
	if todo.ID == 0 {
		errorz.WriteHttpError(w, http.StatusBadRequest, "todo not exist")
		return
	}
	db.Delete(&todo)
	fmt.Fprint(w, "todo was deleted.")

	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}

// title/desc/userid
func AddTodo(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserformRequest(w)
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("something went wrong", err.Error())
		errorz.WriteHttpError(w, http.StatusBadRequest)
		return
	}
	db := connection.GetDB()
	db.Table("todos").Create(&models.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      true,
		UserID:      user.ID,
	})
	fmt.Fprint(w, "Todo add susseccfully. named: ", todo.Title)
	err = db.Close()
	if err != nil {
		log.Println(err)
	}
}

func ViewAll(w http.ResponseWriter, r *http.Request) {

	var allTodo []models.Todo
	db := connection.GetDB()
	db.Find(&allTodo)
	fmt.Fprint(w, allTodo)
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {

	user := auth.GetUserformRequest(w)
	var todos []models.Todo
	db := connection.GetDB()
	var dbAccount models.User
	db.Table("users").Where("user_name =?", user.UserName).First(&dbAccount)
	db.Table("todos").Where("user_id =?", dbAccount.ID).Find(&todos)
	fmt.Fprint(w, todos)
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}
