package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hoorinaz/todo-api/pkg/todo"
	"github.com/hoorinaz/todo-api/pkg/todo/todoservice"
	"github.com/hoorinaz/todo-api/pkg/user/userservice"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/hoorinaz/todo-api/pkg/user"
	usergrpc "github.com/hoorinaz/todo-api/proto/usergrpc"
	"github.com/hoorinaz/todo-api/proto/usergrpc/grpcserver"
	"github.com/hoorinaz/todo-api/shared/connection"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
)

func main() {

	db := connection.GetDB()
	db.AutoMigrate(&todo.Todo{})
	db.AutoMigrate(&user.User{})
	db.Model(&todo.Todo{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	router := mux.NewRouter()

	userservice.RegisterUserService(router)
	todoservice.RegisterTodoService(router)

	fmt.Println("connect to db")

	go http.ListenAndServe(":8080", router)

	///////////////////////////////// grpc server implementation
	log.Println("server starting...")
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Println("failed to listen")
		return
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	userServer := grpcserver.NewGrpcServer()

	usergrpc.RegisterUserServiceServer(server, &userServer)

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic(err.Error())
		}
	}()

	sig := WaitSignal()

	//Blocking...
	log.Println("received signal " + sig.String() + ", exiting...")

	lis.Close()
	server.Stop()
}

func WaitSignal() os.Signal {
	sig := make(chan os.Signal, 1)
	defer close(sig)

	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	return <-sig
}
