package models

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"

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

func AddUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	email := params["email"]

	db := Getdb()
	db.Create(&Account{
		UserName:     username,
		Email:        email,
		Password:     "1234",
		IsRegistered: false,
	})
	fmt.Fprint(w, "add account")
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	hashPass, err := HashPassword(account.Password)
	if err != nil {
		fmt.Fprint(w, "password is clear!")
	}

	db := Getdb()
	db.Create(&Account{
		UserName:     account.UserName,
		Email:        account.Email,
		Password:     hashPass,
		IsRegistered: false,
	})
	fmt.Fprint(w, "add account %v successfully", account.UserName)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		fmt.Fprint(w, "Login Faild")
		fmt.Println(err)
		return
	}
	db := Getdb()
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

	fmt.Fprint(w, "Login Successfuly!")

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
