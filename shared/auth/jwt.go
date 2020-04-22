package auth

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/hoorinaz/TodoList/models"
	"github.com/hoorinaz/TodoList/shared"
	"github.com/hoorinaz/TodoList/shared/store"
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
	JwtKey = []byte("The_Secret_Key")
	//unauthorizedErr = errors.New("Unauthorizaed")
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

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			log.Println("token is not valid ", err.Error())
			w.WriteHeader(http.StatusUnauthorized)

			return
		}
		if !token.Valid {
			log.Println("token is invalid ")
			w.WriteHeader(http.StatusUnauthorized)

			return
		} else {
			dbUser := models.User{}
			db := store.GetDB()
			if err = db.Table("users").Where("user_name =?", claims.Username).First(&dbUser).Error; err != nil {
				log.Println("User Not Found ", err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			uJson, err := json.Marshal(dbUser)
			if err != nil {
				log.Println("something went wrong in marshaling user ", err.Error())
			}
			w.Header().Set(shared.UserFieldInHttpHeader, string(uJson))
			next(w, r)
		}

	}

}
