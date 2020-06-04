package userservice

import (
	"context"
	"log"

	"github.com/hoorinaz/TodoList/shared/jwt"

	"github.com/hoorinaz/TodoList/pkg/user"
	"golang.org/x/crypto/bcrypt"
)

const logger = "user-process"

type UserProcessor struct {
	UserStore UserStoreInterface
	jwt       jwt.JwtProvider
}

type UserProcessorInterface interface {
	SignUp(context.Context, *user.User) error
	SignIn(context.Context, *user.User) error
}

func (up UserProcessor) SignUp(ctx context.Context, u *user.User) error {

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

func (up UserProcessor) SignIn(ctx context.Context, u *user.User) error {
	dbUser := user.User{
		Username: u.Username,
	}

	err := up.UserStore.GetUser(ctx, &dbUser)
	if err != nil {
		log.Println(logger, "error in store layer")
		return err
	}
	if dbUser.Username == "" {
		log.Println(logger, "user not foud")
		return err // error 401 unauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password))
	if err != nil {
		log.Println("password is incorrect")
		return err
	}
	d := jwt.Data{
		Username: u.Username,
		Email:    u.Email,
	}
	log.Println(d)
	tokenString, err := up.jwt.GenerateToken(d)
	if err != nil {
		log.Println(logger, "error in generate token")
		return err
	}
	ctx = context.WithValue(ctx, "Authorization", tokenString)

	return nil
}

func NewUserProcessor(userStore UserStoreInterface) UserProcessorInterface {
	if userStore == nil {
		userStore = NewUserStore()
	}
	return UserProcessor{

		UserStore: userStore,
		jwt:       jwt.NewJwtProvider(),
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
