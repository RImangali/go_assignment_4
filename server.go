package main

import (
    "context"
	"fmt"
	"log"
	"net"
    
	"google.golang.org/grpc"

	pb "go_assignment_4/user"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}

func (s *userService) AddUser(ctx context.Context, user *pb.User) (*pb.UserID, error) {

	return &pb.UserID{Id: 1}, nil
}

func (s *userService) GetUser(ctx context.Context, userID *pb.UserID) (*pb.User, error) {
	return &pb.User{Id: userID.Id, Name: "John Doe", Email: "john.doe@example.com"}, nil
}

func (s *userService) ListUsers(ctx context.Context, empty *pb.Empty) (*pb.UserList, error) {

	users := []*pb.User{
		{Id: 1, Name: "John Doe", Email: "john.doe@example.com"},
		{Id: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
	}
	return &pb.UserList{Users: users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &userService{})
	fmt.Println("gRPC server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}