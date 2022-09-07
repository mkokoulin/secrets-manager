package handlers

import (
	"context"

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

func (gu *GRPCUsers) CreateUser(ctx context.Context, in *pb.RegistrationRequiest) (*pb.RegistrationResponse, error) {
	user := models.User {
		Login: in.Login,
		Password: in.Password,
	}

	err := gu.userService.CreateUser(ctx, user)
	if err != nil {
		return &pb.RegistrationResponse{
			Status: 0,
		}, nil
	}

	token, err := gu.userService.LoginUser(ctx, user)
	if err != nil {
		return &pb.RegistrationResponse{
			Status: 0,
		}, nil
	}

	return &pb.RegistrationResponse{
		Status: 1,
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}