package todo

import (
	"time"

	"github.com/hoorinaz/todo-api/pkg/user"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	IsDone      bool      `json:"is_done"`
	User        user.User `gorm:"foreignkey:UserID"`
	UserID      uint
}
