package todo

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hoorinaz/TodoList/models"
	"github.com/hoorinaz/TodoList/shared"
	"github.com/hoorinaz/TodoList/shared/store"
	"log"
	"net/http"
)

// id
func ViewTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	params := mux.Vars(r)
	Id := params["id"]
	db := store.GetDB()
	db.First(&todo, Id)
	fmt.Println("FOUND: ", todo)
	json.NewEncoder(w).Encode(todo)
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}

//id => param ,title, description  => json
func EditTodo(w http.ResponseWriter, r *http.Request) {
	u := []byte(w.Header().Get(shared.UserFieldInHttpHeader))
	user := models.User{}
	if err := json.Unmarshal(u, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error in un marshaling is: ", err.Error())
		fmt.Fprint(w, "user not found")

		return
	}
	params := mux.Vars(r)
	id := params["id"]
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("bad request error is: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := store.GetDB()
	dbUser := models.User{}
	err = db.Table("todos").Where("Id =?", id).Find(&dbUser).Error
	if err != nil {
		fmt.Fprintf(w, "todo with %v id not found", id)
		log.Printf("problme found to get todo with %v. error: %v", id, err.Error())
		return
	}
	if dbUser.ID == 0 {
		fmt.Fprintf(w, "incorrect ID, todo not found")
		return
	}

	if dbUser.UserName != user.UserName {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "you are unauthorized!")
		return
	}
	if res := db.Model(&todo).
		Where("id = ?", dbUser.ID).Updates(models.Todo{Title: todo.Title,
		Description: todo.Description}); res.Error != nil {
		fmt.Println("error", res.Error.Error())
	}

	fmt.Fprintf(w, "Todo %v has been changed successfully.", id)

	err = db.Close()
	if err != nil {
		log.Println("error in DB session closing ", err.Error())
	}
}

//id
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	u := []byte(w.Header().Get(shared.UserFieldInHttpHeader))
	user := models.User{}
	if err := json.Unmarshal(u, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "user not found")
		log.Println("error in un marshaling is: ", err.Error())

	}

	var todo models.Todo
	params := mux.Vars(r)
	Id := params["Id"]
	db := store.GetDB()
	if res := db.Table("todos").Where("Id =?", Id).Find(&todo); res.Error != nil {
		fmt.Fprintf(w, "connection problem")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(todo)
	if todo.UserID != user.ID {
		fmt.Fprint(w, "you are not authorized")
		return
	}
	if todo.ID == 0 {
		fmt.Fprint(w, "todo not exist")
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
	user := models.User{}
	u := []byte(r.Header.Get(shared.UserFieldInHttpHeader))
	if err := json.Unmarshal(u, &user); err != nil {
		log.Println("error unmarshalling user", err.Error())
		fmt.Fprint(w, "user not found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("something went wrong", err.Error())
		return
	}
	db := store.GetDB()
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
	db := store.GetDB()
	db.Find(&allTodo)
	fmt.Fprint(w, allTodo)
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {

	var user models.User
	var todos []models.Todo
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprint(w, "Login Faild")
		return
	}
	db := store.GetDB()
	var dbAccount models.User
	db.Table("accounts").Where("user_name =?", user.UserName).First(&dbAccount)
	db.Table("todos").Where("user_id =?", dbAccount.ID).Find(&todos)
	fmt.Fprint(w, todos)
	err = db.Close()
	if err != nil {
		log.Println(err)
	}
}
