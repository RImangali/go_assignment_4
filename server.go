package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"

    pb "github.com/go_assignment_4"
)

type server struct {
    users []*pb.User
}

func (s *server) AddUser(ctx context.Context, user *pb.User) (*pb.UserID, error) {
    s.users = append(s.users, user)
    return &pb.UserID{Id: user.Id}, nil
}

func (s *server) GetUser(ctx context.Context, userID *pb.UserID) (*pb.User, error) {
    for _, user := range s.users {
        if user.Id == userID.Id {
            return user, nil
        }
    }
    return nil, grpc.Errorf(grpc.Code(grpc.NotFound), "User not found")
}

func (s *server) ListUsers(empty *pb.Empty, stream pb.UserService_ListUsersServer) error {
    for _, user := range s.users {
        if err := stream.Send(user); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &server{})
    log.Println("Server started listening on port :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
