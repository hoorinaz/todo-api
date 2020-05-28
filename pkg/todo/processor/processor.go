package processor

import (
	"context"

	"github.com/hoorinaz/TodoList/pkg/todo"
)

type TodoProcessor struct {
	TodoStore todo.TodoService
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

func NewTodoProcessor(todoStore todo.TodoService) todo.TodoService {
	return TodoProcessor{

		TodoStore: todoStore,
	}

}
