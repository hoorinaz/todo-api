package user

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hoorinaz/TodoList/todo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Claims struct {
	Username string
	Email    string
	jwt.StandardClaims
}

var jwtKey = []byte("The_Secret_Key")

func AddUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	email := params["email"]

	db := todo.Getdb()
	defer db.Close()
	db.Create(&User{
		UserName:     username,
		Email:        email,
		Password:     "1234",
		IsRegistered: false,
	})
	fmt.Fprint(w, "add account")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	if len(user.Password) < 6 {

		fmt.Fprint(w, "Password is required with more than 6 characters length")
		return

	}
	hashPass, err := HashPassword(user.Password)
	if err != nil {
		fmt.Fprint(w, "password is clear!")
	}

	db := todo.Getdb()
	defer db.Close()
	db.Create(&User{
		UserName:     user.UserName,
		Email:        user.Email,
		Password:     hashPass,
		IsRegistered: false,
	})
	fmt.Fprint(w, "add user successfully ", user.UserName)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var user User
	expTime := time.Now().Add(24 * time.Hour)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprint(w, "Login Faild")
		fmt.Println(err)
		return
	}

	db := todo.Getdb()
	defer db.Close()
	var dbAccount User
	db.Table("accounts").Where("user_name =?", user.UserName).First(&dbAccount)

	if dbAccount.UserName == "" {

		fmt.Fprint(w, "user not exist")
		return
	}
	match := CheckPasswordHash(user.Password, dbAccount.Password)
	if match == false {
		fmt.Fprint(w, "Incorrect Password")
		return

	}
	claims := Claims{
		Username: user.UserName,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.Header().Set("Authorization", stringToken)
	fmt.Fprint(w, "Login Successfuly! Token is ", stringToken)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {

	claims := &Claims{}

	cookieTokenString := r.Header.Get("Authorization")
	if cookieTokenString == "" {
		fmt.Fprint(w, "Unauthorized, Please signin first")
		return
	}

	fmt.Println("cookie token string:", cookieTokenString)
	token, err := jwt.ParseWithClaims(cookieTokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
		}
		log.Println(err)
		return
	}

	fmt.Println(token)
	var user User
	var todos []todo.Todo
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprint(w, "Login Faild")
		return
	}
	db := todo.Getdb()
	defer db.Close()
	var dbAccount User
	db.Table("accounts").Where("user_name =?", user.UserName).First(&dbAccount)
	db.Table("todos").Where("user_id =?", dbAccount.ID).Find(&todos)
	fmt.Fprint(w, todos)

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
