package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/daikideal/go-grpc-sample/user"
	"google.golang.org/grpc"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

// ユーザー一覧を返す
func (s *UserService) ListUsers(ctx context.Context, message *user.ListUsersRequest) (*user.ListUsersResponse, error) {
	res := &user.ListUsersResponse{
		Users: []*user.User{
			{Id: 1, Name: "George"},
			{Id: 2, Name: "Fred"},
			{Id: 3, Name: "Steve"},
		},
	}

	return res, nil
}

func main() {
	port := 8080
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err.Error())
		return
	}

	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &UserService{})
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
