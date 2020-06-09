package gateway

import (
	"context"
	"log"

	"github.com/hoorinaz/todo-api/pkg/user"

	"github.com/hoorinaz/todo-api/pkg/user/userservice"
	"github.com/hoorinaz/todo-api/shared"
	"github.com/hoorinaz/todo-api/shared/errorz"
	jwt2 "github.com/hoorinaz/todo-api/shared/jwt"

	"net/http"
)

type Authentication struct {
	userProcessor userservice.UserProcessorInterface
	jwtProvider   jwt2.JwtProvider
}

type AuthenticationProvider interface {
	AuthMidd(next http.HandlerFunc) http.HandlerFunc
}

func (auth *Authentication) AuthMidd(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		data, err := auth.jwtProvider.DecodeToken(tokenString)
		if err != nil {
			log.Println("Decode token went wrong", err.Error())
			return
		}
		ctx := r.Context()
		dbUser := user.User{}
		dbUser.Username = data.Username

		err = auth.userProcessor.GetUser(ctx, &dbUser)
		if err != nil {
			log.Println("user not found", err)
			errorz.WriteHttpError(w, http.StatusUnauthorized)
			return
		}
		ctx = context.WithValue(ctx, shared.UserInContext, dbUser.ID)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
func NewMiddleware() AuthenticationProvider {
	return &Authentication{
		jwtProvider:   jwt2.NewJwtProvider(),
		userProcessor: userservice.NewUserProcessor(nil),
	}
}
