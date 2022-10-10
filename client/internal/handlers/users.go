// Package handlers includes methods for handling users cases
package handlers

import (
	"context"
	"fmt"

	"github.com/mkokoulin/secrets-manager.git/client/internal/services"
)

// NewUserHandler function for creates new user handler
func NewUserHandler(userClient *services.UserClient) *UserHandler {
	return &UserHandler{
		UserClient: userClient,
	}
}

// UserHandler structure for handling users
type UserHandler struct {
	UserClient *services.UserClient
}

// RegisterUser method for registering a user
func (uh *UserHandler) RegisterUser(ctx context.Context) (string, string, error) {
	login, password, err := uh.getUserCredentials()
	if err != nil {
		return "", "", err
	}
	return uh.UserClient.Register(ctx, login, password)
}

// AuthUser method for auth a user
func (uh *UserHandler) AuthUser(ctx context.Context) (string, string, error) {
	login, password, err := uh.getUserCredentials()
	if err != nil {
		return "", "", err
	}
	return uh.UserClient.Auth(ctx, login, password)
}

func (uh *UserHandler) getUserCredentials() (string, string, error) {
	var login, password string
	fmt.Println("Enter login:")
	_, err := fmt.Scan(&login)
	if err != nil {
		fmt.Println("Error with parse login")
		return "", "", err
	}
	fmt.Println("Enter password:")
	_, err = fmt.Scan(&password)
	if err != nil {
		fmt.Println("Error with parse login")
		return "", "", err
	}
	return login, password, nil
}