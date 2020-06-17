package grpcserver

import (
	"context"
	"log"

	"github.com/hoorinaz/todo-api/pkg/user"
	"github.com/hoorinaz/todo-api/pkg/user/userservice"
	userProto "github.com/hoorinaz/todo-api/proto/user"
)

//GrpcServer GrpcServer
type GrpcServer struct {
	db userservice.UserProcessorInterface
}

//NewGrpcServer NewGrpcServer
func NewGrpcServer() GrpcServer {
	return GrpcServer{
		db: userservice.NewUserProcessor(nil),
	}
}

//GetUser GetUser
func (server *GrpcServer) GetUser(ctx context.Context, r *userProto.Request) (*userProto.User, error) {
	u := &user.User{
		Username: r.GetUsername(),
	}
	err := server.db.GetUser(ctx, u)
	if err != nil {
		log.Println("grpcServer: ", "an error occured in server.db.GetUser()", err)
		return nil, err
	}
	return convertProtoUser(*u), nil
}

func convertProtoUser(u user.User) *userProto.User {
	return &userProto.User{
		ID:       uint32(u.ID),
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
