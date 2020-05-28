package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/shared/errorz"
	"github.com/hoorinaz/TodoList/shared/jwt"
	"golang.org/x/crypto/bcrypt"
)

const logger = "user-web"

type UserWebService struct {
	UserProcessor user.UserService
	jwt           jwt.JwtProvider
}

func (uws UserWebService) Register(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	u := new(user.User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		log.Println(logger, "json decode error is: ", err.Error())
		return
	}

	if err = uws.UserProcessor.AddUser(ctx, u); err != nil {
		log.Println(logger, "error in processor layer", err.Error())
		errorz.WriteHttpError(w, http.StatusInternalServerError)
		return
	}

}

func (uws UserWebService) Authenticate(w http.ResponseWriter, r *http.Request) {
	u := new(user.User)

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		log.Println(logger, "json decode error is: ", err.Error())
		return
	}

	ctx := r.Context()
	dbUser := user.User{
		Username: u.Username,
	}
	err = uws.UserProcessor.GetUser(ctx, &dbUser)

	if err != nil {
		log.Println(logger, "error in processor layer", err)
		errorz.WriteHttpError(w, http.StatusInternalServerError)
		return
	}
	if dbUser.Username == "" {
		errorz.WriteHttpError(w, http.StatusInternalServerError, "user not found")
		log.Println("the user does not exist in DB")
		return
	}

	match := CheckPasswordHash(u.Password, dbUser.Password)
	if match == false {
		errorz.WriteHttpError(w, http.StatusUnauthorized, "incorrect password")
		log.Println("Password is not match")
		return
	}
	d := jwt.Data{
		Username: u.Username,
		Email:    u.Email,
	}

	tokenString, err := uws.jwt.GenerateToken(d)
	if err != nil {
		log.Println("error in tokne generate: ", err)
		return
	}
	// ctx = context.WithValue(ctx, "Authorization", tokenString)
	w.Header().Set("Authorization", tokenString)
	fmt.Fprint(w, tokenString)

}

func NewUserWebService(userProcessor user.UserService) UserWebService {

	return UserWebService{
		UserProcessor: userProcessor,
		jwt:           jwt.NewJwtProvider(),
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
