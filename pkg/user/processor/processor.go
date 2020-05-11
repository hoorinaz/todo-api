package processor

import (
		"github.com/hoorinaz/TodoList/pkg/user"

)
type UserProcessor struct {

	UserStore user.UserService
}

func (up UserProcessor) AddUser(u *user.User) error{


	return up.UserStore.AddUser(u)
}
func(up UserProcessor) GetUser (u *user.User)error{
	return up.UserStore.GetUser(u)
}

func NewUserProcessor (userStore user.UserService) user.UserService{
	return UserProcessor{
		UserStore: userStore,
	}
}

