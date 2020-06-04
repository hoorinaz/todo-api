package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/hoorinaz/TodoList/pkg/todo"
	"github.com/hoorinaz/TodoList/shared"
	"github.com/hoorinaz/TodoList/shared/errorz"
)

const logger = "todo-web"

type TodoWebService struct {
	TodoProcessor todo.TodoService
}

//title, description.....>json body
func (tws TodoWebService) AddTodo(w http.ResponseWriter, r *http.Request) {

	var todo *todo.Todo
	ctx := r.Context()
	userId := r.Header.Get(shared.UserFieldInHttpHeader)

	u32, err := strconv.ParseUint(userId, 10, 32)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		log.Println(logger, "json decode error is: ", err.Error())
		return
	}
	todo.UserID = uint(u32)

	if err = tws.TodoProcessor.AddTodo(ctx, *todo); err != nil {
		log.Println(logger, "error in processor layer", err.Error())
		errorz.WriteHttpError(w, http.StatusInternalServerError)
		return
	}

}

//id ....> json body
func (tws TodoWebService) ViewTodo(w http.ResponseWriter, r *http.Request) {
	var todo *todo.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "bad request")
		return
	}
	ctx := r.Context()
	if err = tws.TodoProcessor.ViewTodo(ctx, todo); err != nil {
		log.Println(logger, "view todo, error in processor layer", err.Error())
		return
	}
	fmt.Fprint(w, todo)
}

func (tws TodoWebService) EditTodo(w http.ResponseWriter, r *http.Request) {
	td := &todo.Todo{}
	userId := r.Header.Get(shared.UserFieldInHttpHeader)
	err := json.NewDecoder(r.Body).Decode(&td)
	if err != nil {
		fmt.Println("bad request", err.Error())
		errorz.WriteHttpError(w, http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	dbTd := &todo.Todo{
		Model: gorm.Model{ID: td.ID},
	}
	if err = tws.TodoProcessor.ViewTodo(ctx, dbTd); err != nil {
		fmt.Println(logger, "get todo got error: ", err.Error())
		return
	}
	u32, err := strconv.ParseUint(userId, 10, 32)
	if dbTd.ID == 0 || dbTd.UserID != uint(u32) {
		fmt.Println(logger, "wrong todo or wrong user ")
		errorz.WriteHttpError(w, http.StatusUnauthorized, "wrong todo or wrong user ")

		return
	}

	if err = tws.TodoProcessor.EditTodo(ctx, td); err != nil {
		fmt.Println(logger, "error in edit todo", err.Error())
	}
	fmt.Fprintf(w, "Todo %v has been changed successfully.", td.ID)

}

func (tws TodoWebService) ListTodo(w http.ResponseWriter, r *http.Request) {
	todos := []todo.Todo{}
	ctx := r.Context()
	tws.TodoProcessor.ListTodo(ctx, &todos)
	fmt.Fprint(w, todos)

}

func (tws TodoWebService) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	td := todo.Todo{}
	err := json.NewDecoder(r.Body).Decode(&td)
	if err != nil {
		log.Println(logger, "bad request")
		return
	}
	userId := r.Header.Get(shared.UserFieldInHttpHeader)
	u32, err := strconv.ParseUint(userId, 10, 32)
	ctx := r.Context()
	dbTodo := todo.Todo{}
	dbTodo.ID = td.ID
	err = tws.TodoProcessor.ViewTodo(ctx, &dbTodo)
	if uint(u32) != dbTodo.UserID || dbTodo.ID == 0 {
		log.Println(logger, "wrog user or worng todo")
		errorz.WriteHttpError(w, http.StatusUnauthorized)
		return
	}
	if err = tws.TodoProcessor.DeleteTodo(ctx, &td); err != nil {
		log.Println(logger, "there is an error in processor layer, ", err.Error())
		return
	}

}

func NewTodoWebService(todoProcessor todo.TodoService) TodoWebService {
	return TodoWebService{
		TodoProcessor: todoProcessor,
	}
}
