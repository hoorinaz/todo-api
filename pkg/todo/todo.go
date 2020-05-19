package todo

import (
	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/jinzhu/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title       string    		 `json:"title"`
	Description string   		 `json:"description"`
	DueDate     time.Time		 `json:"due_date"`
	IsDone      bool      		 `json:"is_done"`
	User        user.User        `gorm:"foreignkey:UserID"`
	UserID      uint
}


type TodoService interface {
	AddTodo(todo Todo) error
	//GetTodo(todo Todo)error
	//EditTodo(todo Todo) error
	//DeleteTodo(todo Todo)error
	//ListTodo(todo Todo)error

}
