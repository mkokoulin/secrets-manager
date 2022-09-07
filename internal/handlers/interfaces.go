package handlers

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/internal/auth"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user models.User) error
	LoginUser(ctx context.Context, user models.User) (*auth.TokenDetails, error)
	DeleteUser(ctx context.Context, userID string) error
	RefreshToken(ctx context.Context, refreshToken string) (*auth.TokenDetails, error)
}