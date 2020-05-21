package processor

import (
	"context"
	"log"

	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/pkg/user/store"
)

type UserProcessor struct {
	UserStore user.UserService
}

func (up *UserProcessor) AddUser(ctx context.Context, u *user.User) error {
	v := ctx.Value("name")
	log.Println("value in the context", v)
	return up.UserStore.AddUser(ctx, u)
}

func (up *UserProcessor) GetUser(ctx context.Context, u *user.User) error {
	return up.UserStore.GetUser(ctx, u)
}

func NewUserProcessor(us user.UserService) user.UserService {
	if us == nil {
		us = store.NewUserStore()
	}

	return &UserProcessor{
		UserStore: us,
	}

}
