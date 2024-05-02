package main

import (
	"context"
	"testing"
	
	pb "go_assignment_4/user"
)

type mockUserServiceClient struct{}

func (c *mockUserServiceClient) AddUser(ctx context.Context, user *pb.User) (*pb.UserID, error) {

	return &pb.UserID{Id: 1}, nil
}

func (c *mockUserServiceClient) GetUser(ctx context.Context, userID *pb.UserID) (*pb.User, error) {
	return &pb.User{Id: userID.Id, Name: "John Doe", Email: "john.doe@example.com"}, nil
}

func (c *mockUserServiceClient) ListUsers(ctx context.Context, empty *pb.Empty) (*pb.UserList, error) {
	users := []*pb.User{
		{Id: 1, Name: "John Doe", Email: "john.doe@example.com"},
		{Id: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
	}
	return &pb.UserList{Users: users}, nil
}

func TestUserServiceClient_AddUser(t *testing.T) {
	client := &mockUserServiceClient{}
	_, err := client.AddUser(context.Background(), &pb.User{})
	if err != nil {
		t.Errorf("AddUser failed: %v", err)
	}
}

func TestUserServiceClient_GetUser(t *testing.T) {
	client := &mockUserServiceClient{}
	_, err := client.GetUser(context.Background(), &pb.UserID{Id: 1})
	if err != nil {
		t.Errorf("GetUser failed: %v", err)
	}
}

func TestUserServiceClient_ListUsers(t *testing.T) {
	client := &mockUserServiceClient{}
	_, err := client.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		t.Errorf("ListUsers failed: %v", err)
	}
}