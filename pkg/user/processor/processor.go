package processor

import (
	"context"
	"log"

	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/pkg/user/store"
	"golang.org/x/crypto/bcrypt"
)

const logger = "user-process"

type UserProcessor struct {
	UserStore user.UserService
}

func (up UserProcessor) AddUser(ctx context.Context, u *user.User) error {

	if len(u.Password) < 6 {
		log.Println(logger, "password is less than 6 character")
		return nil
	}
	hashPass, err := HashPassword(u.Password)
	if err != nil {
		log.Println(logger, "Error in Password Hashing: ", err.Error())
		return nil

	}
	u.Password = hashPass
	return up.UserStore.AddUser(ctx, u)
}

func (up UserProcessor) GetUser(ctx context.Context, u *user.User) error {

	return up.UserStore.GetUser(ctx, u)
}

func NewUserProcessor(userStore user.UserService) user.UserService {
	if userStore == nil {
		userStore = store.NewUserStore()
	}
	return UserProcessor{
		UserStore: userStore,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
