package main

import (
	"context"
	"log"
	"net"

	"backathon/domain"
	"backathon/internal/application"
	"backathon/internal/infrastructure/repository/memory"
	"backathon/proto/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	userUsecase application.UserUsecaseInterface
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := s.userUsecase.GetUser(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}, nil
}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := domain.User{
		Id:        0,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  "",
	}
	newUser, err := s.userUsecase.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:        newUser.Id,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
	}, nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	user := domain.User{
		Id:        in.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
	}
	updatedUser, err := s.userUsecase.UpdateUser(in.Id, user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:        updatedUser.Id,
		FirstName: updatedUser.FirstName,
		LastName:  updatedUser.LastName,
		Email:     updatedUser.Email,
		Password:  updatedUser.Password,
	}, nil
}

func (s *server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.UserResponse, error) {
	err := s.userUsecase.DeleteUser(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{}, nil
}

func main() {
	userRepository := memory.NewUserRepository()

	userUsecase := application.NewUserUsecase(userRepository)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{
		userUsecase: userUsecase,
	})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
