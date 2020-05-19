package apiGetway

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hoorinaz/TodoList/pkg/user/processor"
	userWeb "github.com/hoorinaz/TodoList/pkg/user/web"
	jwt2 "github.com/hoorinaz/TodoList/shared/jwt"
	"log"
	"net/http"
)

func (auth *Authentication) getUser(r *http.Request) error {
	tokenString := r.Header.Get("Authorization")
	claims := &jwt2.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwt2.Key, nil
	})
	if err != nil {
		log.Println("token is not valid")
		return err
	}
	if !token.Valid {
		log.Println("token is invalid")
		return ErrTokenInvalid
	}
	return nil

}

func NewMiddleWare() AuthenticationProvider {
	return &Authentication{
		userService: processor.NewUserProcessor(),
	}
}
