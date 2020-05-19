package web

import (
	"encoding/json"
	"github.com/hoorinaz/TodoList/pkg/todo"
	"github.com/hoorinaz/TodoList/shared/errorz"
	"log"
	"net/http"
)
const logger ="todo-web"

type TodoWebService struct {
	TodoProcessor  todo.TodoService
}

func (tws TodoWebService)AddTodo(w http.ResponseWriter, r *http.Request){
	var todo *todo.Todo
	err:=json.NewDecoder(r.Body).Decode(&todo)
	if err!=nil{
		errorz.WriteHttpError(w, http.StatusBadRequest,"Bad Request")
		log.Println(logger,"json decode error is: ", err.Error())
		return
	}
	if err= tws.TodoProcessor.AddTodo(*todo) ;err!=nil{
		log.Println(logger,"error in processor layer", err.Error())
		errorz.WriteHttpError(w,http.StatusInternalServerError)
		return
	}

}

func NewTodoWebService(todoProcessor todo.TodoService) TodoWebService{
	return TodoWebService{
		TodoProcessor: todoProcessor,
	}
}