package apiGetway

import (
	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/shared/jwt"
	"net/http"
)

type Authentication struct {
	userService user.UserService
	jwtProvider jwt.JwtProvider
}

type AuthenticationProvider interface {
	Au(next http.HandlerFunc) http.HandlerFunc
}
