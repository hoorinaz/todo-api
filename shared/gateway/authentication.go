package gateway

import (
	"context"
	"log"

	// "google.golang.org/grpc"

	"net/http"

	"github.com/hoorinaz/todo-api/pkg/user/userservice"
	userProto "github.com/hoorinaz/todo-api/proto/user"
	"github.com/hoorinaz/todo-api/shared"
	jwt2 "github.com/hoorinaz/todo-api/shared/jwt"
	"google.golang.org/grpc"
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
		// dbUser := user.User{}
		// dbUser.Username = data.Username

		// err = auth.userProcessor.GetUser(ctx, &dbUser)
		// if err != nil {
		// 	log.Println("user not found", err)
		// 	errorz.WriteHttpError(w, http.StatusUnauthorized)
		// 	return
		// }

		//////////////////// grpc client implementation//////////////////////////////

		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			log.Println("server dialing got error", err.Error)
		}
		defer conn.Close()
		client := userProto.NewUserServiceClient(conn)

		u, err := client.GetUser(ctx, &userProto.Request{Username: data.Username})

		ctx = context.WithValue(ctx, shared.UserInContext, u.ID)
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
