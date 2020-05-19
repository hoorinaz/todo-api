package processor

import "github.com/hoorinaz/TodoList/pkg/todo"

type TodoProcessor struct {
	TodoStore  todo.TodoService
}

func (tp TodoProcessor) AddTodo(t todo.Todo) error{

	return tp.TodoStore.AddTodo(t)


}

func NewTodoProcessor(todoStore todo.TodoService)todo.TodoService{
	return TodoProcessor{
		TodoStore: todoStore,
	}

}
