package user

import (
	"encoding/json"
	"fmt"
	"github.com/hoorinaz/TodoList/models"
	"github.com/hoorinaz/TodoList/shared/auth"
	"github.com/hoorinaz/TodoList/shared/store"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error: ", err)
		return
	}
	if len(user.Password) < 6 {
		fmt.Fprint(w, "Password is required with more than 6 characters length")
		return
	}
	hashPass, err := HashPassword(user.Password)
	if err != nil {
		fmt.Fprint(w, "password is clear!")
		log.Println("Error in Password Hashing: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := store.GetDB()
	db.Create(&models.User{
		UserName:     user.UserName,
		Email:        user.Email,
		Password:     hashPass,
		IsRegistered: false,
	})
	fmt.Fprintf(w, "User %v added successfully ", user.UserName)
	stringToken, err := auth.CreateToken(user.UserName, user.Email)
	if err != nil {
		log.Println("create token failed ", err.Error())

	} else {
		w.Header().Set("Authorization", stringToken)
	}
	err = db.Close()
	if err != nil {
		log.Println("Error in  database session closing: ", err.Error())
	}
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprint(w, "Login Faild")
		log.Println("Error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db := store.GetDB()
	var dbUser models.User
	err = db.Table("users").Where("user_name =?", user.UserName).First(&dbUser).Error
	if err != nil {
		fmt.Fprint(w, "something went wrong")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("there is problem to connect DB")
		return

	}
	if dbUser.UserName == "" {
		fmt.Fprint(w, "user not exist")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("the user does not exist in DB")
		return
	}
	match := CheckPasswordHash(user.Password, dbUser.Password)
	if match == false {
		fmt.Fprint(w, "Incorrect Password")
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Password is not match")
		return
	}
	stringToken, err := auth.CreateToken(user.UserName, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Authorization", stringToken)
	fmt.Fprint(w, "login successfully")
	err = db.Close()
	if err != nil {
		log.Println(err)
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
