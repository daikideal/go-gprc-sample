package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/daikideal/go-grpc-sample/pb"
)

type UserService struct{
	pb.UnimplementedUserServiceServer
}

var fixtures []*pb.User

func init() {
	createFixtures()
}

// ユーザー一覧を返す
func (s *UserService) ListUsers(ctx context.Context, message *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	if len(fixtures) < 1 {
		return nil, errors.New("Users not found. Maybe fixtures are not loaded.")
	}

	return &pb.ListUsersResponse{
		Users: fixtures,
	}, nil
}

// ユーザー単体を返す
func (s *UserService) GetUser(ctx context.Context, message *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	q := message.Name
	for _, user := range fixtures {
		if user.Name == q {
			return &pb.GetUserResponse{
				User: user,
			}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("User Name:=%s is not found.", q))
}

// NOTE: ひとまずDBを使用せずに実装するため、メモリ上にユーザーデータを作成する
func createFixtures() {
	fixtures = []*pb.User{
		{Id: 1, Name: "George"},
		{Id: 2, Name: "Fred"},
		{Id: 3, Name: "Willium"},
	}
}
