package userservice

import (
	"context"
	"fmt"
	"log"

	"github.com/hoorinaz/todo-api/shared/jwt"

	"github.com/hoorinaz/todo-api/pkg/user"
	"golang.org/x/crypto/bcrypt"
)

const logger = "user-process"

type UserProcessor struct {
	UserStore UserStoreInterface
	jwt       jwt.JwtProvider
}

type UserProcessorInterface interface {
	Register(context.Context, *user.User) error
	Authenticate(context.Context, *user.User) error
}

func (up UserProcessor) Register(ctx context.Context, u *user.User) error {

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

func (up UserProcessor) Authenticate(ctx context.Context, u *user.User) error {
	dbUser := user.User{
		Username: u.Username,
	}

	err := up.UserStore.GetUser(ctx, &dbUser)
	fmt.Println("dbUser  form user processor, ", dbUser.Password)
	fmt.Println("U  form user processor, ", u.Password)

	if err != nil {
		log.Println(logger, "error in store layer ", err.Error())
		return err
	}
	if dbUser.Username == "" {
		log.Println(logger, "user not foud, ", err.Error())
		return err // error 401 unauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password))
	if err != nil {
		log.Println("password is incorrect: ", err.Error())
		return err
	}
	d := jwt.Data{
		Username: u.Username,
	}
	tokenString, err := up.jwt.GenerateToken(d)
	if err != nil {
		log.Println(logger, "error in generate token", err.Error())
		return err
	}
	u.Token = tokenString
	// ctx = context.WithValue(ctx, "Authorization", tokenString)
	log.Println("user processor set context: ", ctx.Value("Authorization"))

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
