package todoservice

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/hoorinaz/todo-api/shared"

	"github.com/hoorinaz/todo-api/pkg/todo"
)

type TodoProcessor struct {
	TodoStore TodoStoreInterface
}

type TodoServiceInterface interface {
	AddTodo(ctx context.Context, todo *todo.Todo) error
	ViewTodo(ctx context.Context, todo *todo.Todo) error
	EditTodo(ctx context.Context, todo *todo.Todo) error
	ListTodo(ctx context.Context, todos *[]todo.Todo) error
	DeleteTodo(ctx context.Context, todo *todo.Todo) error
}

func (tp TodoProcessor) AddTodo(ctx context.Context, t *todo.Todo) error {

	return tp.TodoStore.AddTodo(ctx, t)

}

func (tp TodoProcessor) ViewTodo(ctx context.Context, t *todo.Todo) error {

	return tp.TodoStore.GetTodo(ctx, t)

}
func (tp TodoProcessor) EditTodo(ctx context.Context, t *todo.Todo) error {

	userID := ctx.Value(shared.UserInContext).(uint)
	dbTodo := todo.Todo{
		Model: gorm.Model{ID: t.ID},
	}

	if err := tp.TodoStore.GetTodo(ctx, &dbTodo); err != nil {
		log.Println(logger, "there is an error in store layer", err.Error())
		return err
	}
	if userID != dbTodo.UserID {
		log.Println(logger, "unauthorized")
		return ErrUnauthorized
	}

	if dbTodo.ID == 0 {
		log.Println(logger, "not found todo with ID= ", t.ID)
		return ErrNotFound
	}

	return tp.TodoStore.EditTodo(ctx, t)
}
func (tp TodoProcessor) ListTodo(ctx context.Context, t *[]todo.Todo) error {

	return tp.TodoStore.ListTodo(ctx, t)
}
func (tp TodoProcessor) DeleteTodo(ctx context.Context, t *todo.Todo) error {

	userID := ctx.Value(shared.UserInContext).(uint)
	dbTodo := todo.Todo{
		Model: gorm.Model{ID: t.ID},
	}
	if userID != dbTodo.UserID {
		log.Println(logger, "Unautorized")
		return ErrUnauthorized
	}

	if dbTodo.ID == 0 {
		log.Println(logger, "not found todo with ID= ", t.ID)
		return ErrNotFound

	}
	return tp.TodoStore.DeleteTodo(ctx, t)
}

func NewTodoProcessor(todoStore TodoStoreInterface) TodoServiceInterface {
	return TodoProcessor{

		TodoStore: todoStore,
	}

}
