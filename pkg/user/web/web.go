package web

import "github.com/hoorinaz/TodoList/pkg/user"

type UserWebService struct {
	UserProcessor user.UserService
}

func NewUserWebService(us user.UserService) *UserWebService {
	return &UserWebService{
		UserProcessor: us,
	}
}
