package grpcserver

import (
	"context"
	"log"

	"github.com/hoorinaz/todo-api/pkg/user"
	"github.com/hoorinaz/todo-api/pkg/user/userservice"
	usergrpc "github.com/hoorinaz/todo-api/proto/usergrpc"
)

//GrpcServer struct
type GrpcServer struct {
	uProcessor userservice.UserProcessorInterface
}

//NewGrpcServer create NewUserProcessor
func NewGrpcServer() GrpcServer {
	return GrpcServer{
		uProcessor: userservice.NewUserProcessor(nil),
	}
}

//GetUser GetUser
func (server *GrpcServer) GetUser(ctx context.Context, r *usergrpc.Request) (*usergrpc.User, error) {
	u := &user.User{
		Username: r.GetUsername(),
	}
	err := server.uProcessor.GetUser(ctx, u)
	if err != nil {
		log.Println("grpcServer: ", "an error occured in server.db.GetUser()", err)
		return nil, err
	}
	return convertProtoUser(*u), nil
}

func convertProtoUser(u user.User) *usergrpc.User {
	return &usergrpc.User{
		ID:       uint32(u.ID),
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
