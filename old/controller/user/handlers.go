package user

import (
	"encoding/json"
	"fmt"
	"github.com/hoorinaz/TodoList/old/models"
	"github.com/hoorinaz/TodoList/shared/auth"
	"github.com/hoorinaz/TodoList/shared/connection"
	"github.com/hoorinaz/TodoList/shared/errorz"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		log.Println("Json Decode error is : ", err)
		return
	}
	if len(user.Password) < 6 {
		fmt.Fprint(w, "Password is required with more than 6 characters length")
		return
	}
	hashPass, err := HashPassword(user.Password)
	if err != nil {
		log.Println("Error in Password Hashing: ", err.Error())
		errorz.WriteHttpError(w, http.StatusBadRequest)
		return
	}

	db := connection.GetDB()
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
		log.Println("Json Decode Error is: ", err.Error())
		errorz.WriteHttpError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	db := connection.GetDB()
	var dbUser models.User
	err = db.Table("users").Where("user_name =?", user.UserName).First(&dbUser).Error
	if err != nil {
		errorz.WriteHttpError(w, http.StatusInternalServerError, "something went wrong")
		log.Println("there is problem to connect DB", err.Error())
		return

	}
	match := CheckPasswordHash(user.Password, dbUser.Password)
	if match == false{
		errorz.WriteHttpError(w, http.StatusUnauthorized, "incorrect password")
		log.Println("Password is not match")
		return
	}
	if dbUser.UserName == "" {
		errorz.WriteHttpError(w, http.StatusInternalServerError, "user not found")
		log.Println("the user does not exist in DB")
		return
	}
	stringToken, err := auth.CreateToken(user.UserName, user.Email)
	if err != nil {
		errorz.WriteHttpError(w, http.StatusInternalServerError)
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
