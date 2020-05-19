package store

import (
	"github.com/hoorinaz/TodoList/pkg/todo"
	"github.com/hoorinaz/TodoList/shared/connection"
	"github.com/jinzhu/gorm"
	"log"
)

type TodoStore struct {
	DB *gorm.DB
}

const logger = "todo-store"
func(ts TodoStore) AddTodo(td todo.Todo)  error{
	db:=ts.DB
	createTodo:= db.Create(&todo.Todo{
		Title: td.Title,
		Description:td.Description,

	})
if createTodo.Error!=nil{
	log.Println(logger,"there is an error to connection: ", createTodo.Error)
	return nil
	}
db.Create(&todo.Todo{
		Title: td.Title,
		Description:td.Description,
		UserID: 5,

	})
return nil
}

func NewTodoStore() todo.TodoService{

	s:= connection.GetDB()
	return TodoStore{
		DB: s,
	}
}