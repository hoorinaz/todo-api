package todo

import "time"

type Todo struct {
	ID          uint      `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	IsDone      bool      `db:"is_done"`
	UserID      uint      `db:"fk_user_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type TodoService interface {
	AddTodo(*Todo) error
	EditTodo(*Todo) error
	GetTodo(*Todo) error
	ListTodo(*Todo, []*Todo) error
}
