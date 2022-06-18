package main

import (
	"context"
	"fmt"
	"log"

	"github.com/daikideal/go-grpc-sample/user"
	"google.golang.org/grpc"
)

func main() {
	port := 8080
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)
	message := &user.ListUsersRequest{}
	res, err := client.ListUsers(context.Background(), message)
	if err != nil {
		log.Fatal("Request failed: ", err)
	}

	fmt.Printf("Result: %v\n", res)
}
