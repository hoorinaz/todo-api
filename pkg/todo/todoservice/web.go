package todoservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hoorinaz/todo-api/pkg/todo"
	"github.com/hoorinaz/todo-api/shared/errorz"
)

const loggerW = "todo-web"

type TodoWebService struct {
	TodoProcessor TodoServiceInterface
}

type TodoWebServiceInterface interface {
	AddTodo(w http.ResponseWriter, r *http.Request)
	ViewTodo(w http.ResponseWriter, r *http.Request)
	EditTodo(w http.ResponseWriter, r *http.Request)
	ListTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

//title, description.....>json body
func (tws TodoWebService) AddTodo(w http.ResponseWriter, r *http.Request) {

	var todo *todo.Todo
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		log.Println(logger, "json decode error is: ", err.Error())
		return
	}

	if err = tws.TodoProcessor.AddTodo(ctx, todo); err != nil {
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
		log.Println(loggerW, err)
		errorz.WriteHttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	fmt.Fprint(w, todo)
}

func (tws TodoWebService) EditTodo(w http.ResponseWriter, r *http.Request) {
	td := &todo.Todo{}
	err := json.NewDecoder(r.Body).Decode(&td)
	if err != nil {
		fmt.Println("bad request", err.Error())
		errorz.WriteHttpError(w, http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	if err = tws.TodoProcessor.EditTodo(ctx, td); err != nil {
		log.Println(loggerW, "error in edit todo", err.Error())
		errorz.WriteHttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	fmt.Fprintf(w, "Todo %v has been changed successfully.", td.ID)

}

func (tws TodoWebService) ListTodo(w http.ResponseWriter, r *http.Request) {
	todos := []todo.Todo{}
	tws.TodoProcessor.ListTodo(r.Context(), &todos)
	fmt.Fprint(w, todos)

}

func (tws TodoWebService) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	td := todo.Todo{}
	err := json.NewDecoder(r.Body).Decode(&td)
	if err != nil {
		log.Println(logger, "bad request")
		return
	}
	ctx := r.Context()
	if err = tws.TodoProcessor.DeleteTodo(ctx, &td); err != nil {
		log.Println(logger, "there is an error in processor layer, ", err.Error())
		errorz.WriteHttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

}

func NewTodoWebService(todoProcessor TodoServiceInterface) TodoWebServiceInterface {
	return TodoWebService{
		TodoProcessor: todoProcessor,
	}
}
