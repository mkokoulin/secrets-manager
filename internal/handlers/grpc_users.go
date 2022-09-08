package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
	pb "github.com/mkokoulin/secrets-manager.git/internal/pb/users"
)

type GRPCUsers struct {
	pb.UnimplementedUsersServer
	userService UserServiceInterface
}

func NewGRPCUsers(userService UserServiceInterface) *GRPCUsers {
	return &GRPCUsers {
		userService: userService,
	}
}

func (gu *GRPCUsers) CreateUser(ctx context.Context, in *pb.CreateUserRequiest) (*pb.CreateUserResponse, error) {
	user := models.User {
		ID: uuid.New(),
		Login: in.Login,
		Password: in.Password,
	}

	err := gu.userService.CreateUser(ctx, user)
	if err != nil {
		return &pb.CreateUserResponse{
			Status: 1,
		}, nil
	}

	token, err := gu.userService.AuthUser(ctx, user)
	if err != nil {
		return &pb.CreateUserResponse{
			Status: 1,
		}, nil
	}

	return &pb.CreateUserResponse{
		Status: 0,
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}