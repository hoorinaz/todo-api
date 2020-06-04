package gateway

import (
	"log"

	"github.com/hoorinaz/TodoList/pkg/user"
	"github.com/hoorinaz/TodoList/pkg/user/userservice"
	"github.com/hoorinaz/TodoList/shared"
	"github.com/hoorinaz/TodoList/shared/errorz"
	jwt2 "github.com/hoorinaz/TodoList/shared/jwt"

	"net/http"
)

type Authentication struct {
	userProcessor userservice.UserProcessor
	jwtProvider   jwt2.JwtProvider
}

type AuthenticationProvider interface {
	Authenticate(next http.HandlerFunc) http.HandlerFunc
}

func (auth *Authentication) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		ctx := r.Context()
		data, err := auth.jwtProvider.DecodeToken(tokenString)
		if err != nil {
			log.Println("Decode token went wrong", err.Error())
			return
		}
		dbUser := user.User{}
		dbUser.Username = data.Username
		err = auth.userService.GetUser(ctx, &dbUser)
		if err != nil {
			log.Println("user not found", err)
			errorz.WriteHttpError(w, http.StatusUnauthorized)
			return

		}
		w.Header().Set(shared.UserFieldInHttpHeader, string(dbUser.ID))
		// fmt.Fprint(w, "dbUser.ID form gateway package: ", dbUser.ID)
		next(w, r)
	}
}
func NewMiddleware() AuthenticationProvider {
	return &Authentication{
		jwtProvider: jwt2.NewJwtProvider(),
		userService: processor.NewUserProcessor(nil),
	}
}
