package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	UserName     string
	Email        string
	Password     string
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

var jwtkey = []byte("my_Secret_Key")

type loginInfo struct {
	Password string `json:"password"`
	UserName string `json:"username"`
}

type claims struct {
	Username string
	jwt.StandardClaims
}

var users = map[string]string{
	"Hoori": "123",
	"Amir":  "456",
}

func Signin(w http.ResponseWriter, r *http.Request) {

	var login loginInfo
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
	}

	pass, ok := users[login.UserName]

	if !ok || pass != login.Password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(ok)
		return
	}

	expTime := time.Now().Add(5 * time.Minute)

	claims := claims{
		Username: login.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(jwtkey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "JWTToken",
		Value:    stringToken,
		Expires:  expTime,
		HttpOnly: true,
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("JWTToken")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var tknString = cookie.Value
	claims := &claims{}

	token, err := jwt.ParseWithClaims(tknString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
		}
		log.Println(err)
		return
	}

	fmt.Println(token)
	fmt.Fprint(w, "Welcome ", claims.Username)
}
