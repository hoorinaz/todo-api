package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hoorinaz/todo-api/pkg/todo"
	"github.com/hoorinaz/todo-api/pkg/todo/todoservice"
	"github.com/hoorinaz/todo-api/pkg/user/userservice"
	userProto "github.com/hoorinaz/todo-api/proto/user"
	"github.com/hoorinaz/todo-api/proto/user/grpcserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	//"github.com/hoorinaz/TodoList/controller/user"
	//"github.com/hoorinaz/TodoList/models"
	"github.com/hoorinaz/todo-api/pkg/user"

	"net/http"

	"github.com/hoorinaz/todo-api/shared/connection"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db := connection.GetDB()
	db.AutoMigrate(&todo.Todo{})
	db.AutoMigrate(&user.User{})
	db.Model(&todo.Todo{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	router := mux.NewRouter()

	userservice.RegisterUserService(router)
	todoservice.AddTodoService(router)

	fmt.Println("connect to db")

	http.ListenAndServe(":8080", router)
	//////////////////////////////////////////////////////////////////server implementation
	grpclog.Println("server starting...")
	lis, err := net.Listen("tpc", ":8080")
	if err != nil {
		log.Println("failed to listen")
		return
	}
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	userServer := grpcserver.NewGrpcServer()
	userProto.RegisterUserServiceServer(server, &userServer)

	err = server.Serve(lis)
	if err != nil {
		log.Println("serve grpc server got error", err.Error)
	}
}
