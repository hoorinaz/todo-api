package gateway

import (
	"context"
	"fmt"
	"log"

	// "google.golang.org/grpc"

	"net/http"

	usergrpc "github.com/hoorinaz/todo-api/proto/usergrpc"
	"github.com/hoorinaz/todo-api/shared"
	jwt2 "github.com/hoorinaz/todo-api/shared/jwt"
	"google.golang.org/grpc"
)

type Authentication struct {
	jwtProvider jwt2.JwtProvider
	userService usergrpc.UserServiceClient
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
		u, err := auth.userService.GetUser(ctx, &usergrpc.Request{Username: data.Username})
		if err != nil {
			log.Println("auth ", "error to GetUser, ", err.Error())
			return
		}
		fmt.Println(u)
		ctx = context.WithValue(ctx, shared.UserInContext, uint(u.ID))
		r = r.WithContext(ctx)
		next(w, r)
	}
}

func NewMiddleware() AuthenticationProvider {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Println("server dialing got error", err.Error())
	}
	// defer conn.Close()
	client := usergrpc.NewUserServiceClient(conn)
	log.Println("server is dialed...")
	return &Authentication{
		jwtProvider: jwt2.NewJwtProvider(),
		userService: client,
	}
}
