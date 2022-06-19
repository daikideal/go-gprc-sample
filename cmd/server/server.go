package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/daikideal/go-grpc-sample/pb"
	"github.com/daikideal/go-grpc-sample/pkg/user"
)

func main() {
	port := 8080
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err.Error())
		return
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &user.UserService{})
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
