package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hoorinaz/TodoList/old/models"
	"github.com/hoorinaz/TodoList/shared/connection"
	"log"
	"net/http"
	"time"
)

type Claims struct {
	Username string
	Email    string
	jwt.StandardClaims
}

var (
	JwtKey          = []byte("The_Secret_Key")
	unauthorizedErr = errors.New("Unauthorizaed")
)
var expTime = time.Now().Add(24 * time.Hour)

func CreateToken(username string, email string) (string, error) {
	claims := Claims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(JwtKey)
	return stringToken, err
}

func Athorization(tokenString string) (*models.User, error) {

	claims := &Claims{}
	//claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})

	if err != nil {
		fmt.Println("token is not valid", err.Error())
		return nil, err
	}
	if !token.Valid {
		fmt.Println("token is invalid")
		return nil, err
	}
	u := &models.User{
		UserName: claims.Username,
	}
	return u, err
}

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := jwtValidation(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			db := connection.GetDB()
			var dbUser models.User
			if err := db.Table(" ").Where("user_name =?", u.UserName).First(&dbUser).Error; err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			uJson, err := json.Marshal(&dbUser)
			if err != nil {
				log.Println("error marshalling user information", err.Error())
				return
			}
			r.Header.Set("user", string(uJson))
			next(w, r)
		}
	}
}

func jwtValidation(r *http.Request) (*models.User, error) {
	claims := &Claims{}
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})

	if err != nil {
		log.Println("token is not valid ", err.Error())
		return nil, unauthorizedErr
	}
	if !token.Valid {
		log.Println("token is invalid")
		return nil, unauthorizedErr
	}
	u := &models.User{
		UserName: claims.Username,
	}
	return u, nil
}
