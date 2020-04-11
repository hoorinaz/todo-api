package models

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	UserName     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsRegistered bool
	Token        string
}

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

	db := Getdb()
	defer db.Close()
	db.Create(&Account{
		UserName:     username,
		Email:        email,
		Password:     "1234",
		IsRegistered: false,
	})
	fmt.Fprint(w, "add account")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	if len(account.Password) < 6 {

		fmt.Fprint(w, "Password is required with more than 6 characters length")
		return

	}
	hashPass, err := HashPassword(account.Password)
	if err != nil {
		fmt.Fprint(w, "password is clear!")
	}

	db := Getdb()
	defer db.Close()
	db.Create(&Account{
		UserName:     account.UserName,
		Email:        account.Email,
		Password:     hashPass,
		IsRegistered: false,
	})
	fmt.Fprint(w, "add account %v successfully", account.UserName)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var account Account
	expTime := time.Now().Add(20 * time.Minute)

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		fmt.Fprint(w, "Login Faild")
		fmt.Println(err)
		return
	}

	claims := Claims{
		Username: account.UserName,
		Email:    account.Email,
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
	db := Getdb()
	defer db.Close()
	var dbAccount Account
	db.Table("accounts").Where("user_name =?", account.UserName).First(&dbAccount)

	if dbAccount.UserName == "" {

		fmt.Fprint(w, "user not exist")
		return
	}
	match := CheckPasswordHash(account.Password, dbAccount.Password)
	if match == false {
		fmt.Fprint(w, "Incorrect Password")
		return

	}
	w.Header().Set("Authorization", stringToken)
	http.SetCookie(w, &http.Cookie{
		Name:    "JWTToken",
		Value:   stringToken,
		Expires: expTime,
	})
	fmt.Fprint(w, "Login Successfuly!")
	fmt.Fprint(w, "Token is ", stringToken)

}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("JWTToken")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	claims := &Claims{}
	cookieTokenString := cookie.Value
	tokenString := w.Header().Get("Authorization")
	if cookieTokenString == "" {
		fmt.Fprint(w, "Unauthorized, Please signin first")
		return
	}
	fmt.Println("token String : ", tokenString)
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
	var account Account
	var todos []Todo
	err = json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		fmt.Fprint(w, "Login Faild")
		return
	}
	db := Getdb()
	defer db.Close()
	var dbAccount Account
	db.Table("accounts").Where("user_name =?", account.UserName).First(&dbAccount)
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
