package todoservice

import (
	"context"

	"github.com/hoorinaz/todo-api/pkg/todo"
)

type TodoProcessor struct {
	TodoStore TodoStoreInterface
}

type TodoServiceInterface interface {
	AddTodo(ctx context.Context, todo todo.Todo) error
	ViewTodo(ctx context.Context, todo *todo.Todo) error
	EditTodo(ctx context.Context, todo *todo.Todo) error
	ListTodo(ctx context.Context, todos *[]todo.Todo) error
	DeleteTodo(ctx context.Context, todo *todo.Todo) error
}

func (tp TodoProcessor) AddTodo(ctx context.Context, t todo.Todo) error {

	return tp.TodoStore.AddTodo(ctx, t)

}

func (tp TodoProcessor) ViewTodo(ctx context.Context, t *todo.Todo) error {

	return tp.TodoStore.ViewTodo(ctx, t)

}
func (tp TodoProcessor) EditTodo(ctx context.Context, t *todo.Todo) error {

	return tp.TodoStore.EditTodo(ctx, t)
}
func (tp TodoProcessor) ListTodo(ctx context.Context, t *[]todo.Todo) error {
	ctx.Value("CtxUserID")

	return tp.TodoStore.ListTodo(ctx, t)
}
func (tp TodoProcessor) DeleteTodo(ctx context.Context, t *todo.Todo) error {

	return tp.TodoStore.DeleteTodo(ctx, t)
}

func NewTodoProcessor(todoStore TodoStoreInterface) TodoServiceInterface {
	return TodoProcessor{

		TodoStore: todoStore,
	}

}
