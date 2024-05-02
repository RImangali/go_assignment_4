package main

import (
	"context"
	"testing"
	
	pb "assignment/user"
)

type mockUserServiceServer struct{}

func (s *mockUserServiceServer) AddUser(ctx context.Context, user *pb.User) (*pb.UserID, error) {
	return &pb.UserID{Id: 1}, nil
}

func (s *mockUserServiceServer) GetUser(ctx context.Context, userID *pb.UserID) (*pb.User, error) {
	return &pb.User{Id: userID.Id, Name: "John Doe", Email: "john.doe@example.com"}, nil
}

func (s *mockUserServiceServer) ListUsers(ctx context.Context, empty *pb.Empty) (*pb.UserList, error) {
	users := []*pb.User{
		{Id: 1, Name: "John Doe", Email: "john.doe@example.com"},
		{Id: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
	}
	return &pb.UserList{Users: users}, nil
}

func TestUserServiceServer_AddUser(t *testing.T) {
	server := &mockUserServiceServer{}
	_, err := server.AddUser(context.Background(), &pb.User{})
	if err != nil {
		t.Errorf("AddUser failed: %v", err)
	}
}

func TestUserServiceServer_GetUser(t *testing.T) {
	server := &mockUserServiceServer{}
	_, err := server.GetUser(context.Background(), &pb.UserID{Id: 1})
	if err != nil {
		t.Errorf("GetUser failed: %v", err)
	}
}

func TestUserServiceServer_ListUsers(t *testing.T) {
	server := &mockUserServiceServer{}
	_, err := server.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		t.Errorf("ListUsers failed: %v", err)
	}
}