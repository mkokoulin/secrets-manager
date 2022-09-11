package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
	pb "github.com/mkokoulin/secrets-manager.git/internal/pb/users"
	customerrors "github.com/mkokoulin/secrets-manager.git/internal/errors"
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
			Status: customerrors.ParseError(err),
		}, nil
	}

	token, err := gu.userService.AuthUser(ctx, user)
	if err != nil {
		return &pb.CreateUserResponse{
			Status: customerrors.ParseError(err),
		}, nil
	}

	return &pb.CreateUserResponse{
		Status: "created",
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (gu *GRPCUsers) AuthUser(ctx context.Context, in *pb.AuthUserRequest) (*pb.AuthUserResponse, error) {
	user := models.User {
		Login: in.Login,
		Password: in.Password,
	}

	token, err := gu.userService.AuthUser(ctx, user)
	if err != nil {
		return &pb.AuthUserResponse{
			Status: customerrors.ParseError(err),
		}, nil
	}

	return &pb.AuthUserResponse {
		Status: "ok",
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, err
}

func (gu *GRPCUsers) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := gu.userService.DeleteUser(ctx, in.Login)
	if err != nil {
		return &pb.DeleteUserResponse{
			Status: customerrors.ParseError(err),
		}, nil
	}

	return &pb.DeleteUserResponse{
		Status: "ok",
	}, nil
}

func (gu *GRPCUsers) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	token, err := gu.userService.RefreshToken(ctx, in.RefreshToken)
	if err != nil {
		return &pb.RefreshTokenResponse{
			Status: customerrors.ParseError(err),
		}, nil
	}

	return &pb.RefreshTokenResponse{
		Status: "ok",
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}