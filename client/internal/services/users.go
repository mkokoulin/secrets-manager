// Package services includes wrappers for pd files and client
package services

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/mkokoulin/secrets-manager.git/client/internal/pb/users"
)

// UserClient structure for user client
type UserClient struct {
	address string
}

// NewUserClient function for creates new user client
func NewUserClient(address string) *UserClient {
	return &UserClient{
		address: address,
	}
}

// GRPCUser structure for GRPC user
type GRPCUser struct {
	pb.UsersClient
	closeFunc func() error
}

// Register methods for registration new users
func (u *UserClient) Register(ctx context.Context, login string, password string) (string, string, error) {
	client, err := u.getConn()
	if err != nil {
		return "", "", err
	}
	message := pb.CreateUserRequiest{
		Login:    login,
		Password: password,
	}
	response, err := client.CreateUser(ctx, &message)
	if err != nil {
		return "", "", err
	}

	if response.Status == "created" {
		return response.AccessToken, response.RefreshToken, nil
	}

	return "", "", errors.New(response.Status)
}

// Auth methods for users authentication
func (u *UserClient) Auth(ctx context.Context, login string, password string) (string, string, error) {
	client, err := u.getConn()
	if err != nil {
		return "", "", err
	}
	message := pb.AuthUserRequest{
		Login:    login,
		Password: password,
	}
	response, err := client.AuthUser(ctx, &message)
	if err != nil {
		return "", "", err
	}
	if response.Status == "ok" {
		return response.AccessToken, response.RefreshToken, nil
	}
	return "", "", errors.New(response.Status)
}

// Refresh method for refreshing auth token
func (u *UserClient) Refresh(ctx context.Context, refreshToken string) (string, string, error) {
	client, err := u.getConn()
	if err != nil {
		return "", "", err
	}
	message := pb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	}
	response, err := client.RefreshToken(ctx, &message)
	if err != nil {
		return "", "", nil
	}
	if response.Status == "ok" {
		return response.AccessToken, response.RefreshToken, nil
	}
	return "", "", errors.New(response.Status)
}

func (u *UserClient) getConn() (*GRPCUser, error) {
	conn, err := grpc.Dial(u.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	cl := pb.NewUsersClient(conn)

	return &GRPCUser{cl, conn.Close}, nil
}
