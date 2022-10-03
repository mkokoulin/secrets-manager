package handlers

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/server/internal/auth"
	"github.com/mkokoulin/secrets-manager.git/server/internal/models"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user models.User) error
	AuthUser(ctx context.Context, user models.User) (*auth.TokenDetails, error)
	DeleteUser(ctx context.Context, userID string) error
	RefreshToken(ctx context.Context, refreshToken string) (*auth.TokenDetails, error)
}

type SecretServiceInterface interface {
	AddSecret(ctx context.Context, secret models.RawSecretData) error
	GetSecret(ctx context.Context, secretID, userID string) (models.Secret, error)
	DeleteSecret(ctx context.Context, secretID, userID string) error
}