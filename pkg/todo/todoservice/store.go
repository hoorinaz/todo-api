package todoservice

import (
	"context"
	"log"

	"github.com/hoorinaz/todo-api/pkg/todo"
	"github.com/hoorinaz/todo-api/shared/connection"
	"github.com/jinzhu/gorm"
)

type TodoStore struct {
	DB *gorm.DB
}
type TodoStoreInterface interface {
	AddTodo(ctx context.Context, todo *todo.Todo) error
	GetTodo(ctx context.Context, todo *todo.Todo) error
	EditTodo(ctx context.Context, todo *todo.Todo) error
	ListTodo(ctx context.Context, todos *[]todo.Todo) error
	DeleteTodo(ctx context.Context, todo *todo.Todo) error
}

const logger = "todo-store"

func (ts TodoStore) AddTodo(ctx context.Context, td *todo.Todo) error {
	db := ts.DB

	db.Create(&todo.Todo{
		Title:       td.Title,
		Description: td.Description,
		UserID:      td.UserID,
	})
	return nil
}

func (ts TodoStore) GetTodo(ctx context.Context, todo *todo.Todo) error {
	db := ts.DB
	if err := db.Table("todos").Where("Id=?", todo.ID).First(&todo).Error; err != nil {
		log.Println(logger, "there is problem to get todo, error: ", todo.ID, " error= ", err.Error())
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

func (ts TodoStore) ListTodo(ctx context.Context, todos *[]todo.Todo) error {
	db := ts.DB
	if err := db.Table("todos").Find(todos).Error; err != nil {
		log.Println(logger, "todo is not found", err)
		return err
	}
	return nil
}

func (ts TodoStore) DeleteTodo(ctx context.Context, todo *todo.Todo) error {
	db := ts.DB
	if err := db.Table("todos").Delete(&todo).Error; err != nil {
		log.Println(logger, err.Error())
		return err

	}
	return nil
}

func NewTodoStore() TodoStoreInterface {

	s := connection.GetDB()
	return TodoStore{
		DB: s,
	}
}
