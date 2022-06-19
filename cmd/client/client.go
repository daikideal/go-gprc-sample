package main

import (
	"context"
	"fmt"
	"log"
	
	"google.golang.org/grpc"

	"github.com/daikideal/go-grpc-sample/pb"
)

func main() {
	port := 8080
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// ユーザー一覧の取得
	msgListUser := &pb.ListUsersRequest{}
	resListUser, err := client.ListUsers(context.Background(), msgListUser)
	if err != nil {
		log.Fatal("Request failed: ", err)
	}
	fmt.Printf("Result: %v\n", resListUser)

	// ユーザー一覧の取得
	msgGetUser := &pb.GetUserRequest{Name: "Fred"}
	resGetUser, err := client.GetUser(context.Background(), msgGetUser)
	if err != nil {
		log.Fatal("Request failed: ", err)
	}
	fmt.Printf("Result: %v\n", resGetUser)
}
