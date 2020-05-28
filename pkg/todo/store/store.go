package store

import (
	"context"
	"log"

	"github.com/hoorinaz/TodoList/pkg/todo"
	"github.com/hoorinaz/TodoList/shared/connection"
	"github.com/jinzhu/gorm"
)

type TodoStore struct {
	DB *gorm.DB
}

const logger = "todo-store"

func (ts TodoStore) AddTodo(ctx context.Context, td todo.Todo) error {
	db := ts.DB

	db.Create(&todo.Todo{
		Title:       td.Title,
		Description: td.Description,
		UserID:      td.UserID,
	})
	return nil
}

func (ts TodoStore) ViewTodo(ctx context.Context, todo *todo.Todo) error {
	db := ts.DB
	if err := db.Table("todos").Where("Id=?", todo.ID).First(&todo).Error; err != nil {
		log.Println(logger, "there is problem to get todo, error: ", err.Error())
		return err
	}
	return nil
}

func (ts TodoStore) EditTodo(ctx context.Context, todo *todo.Todo) error {
	db := ts.DB
	if err := db.Table("todos").Where("Id=?", todo.ID).Update(&todo).Error; err != nil {
		log.Println(logger, "todo is not found", err.Error())
		return err
	}

	return nil
}

func NewTodoStore() todo.TodoService {

	s := connection.GetDB()
	return TodoStore{
		DB: s,
	}
}
