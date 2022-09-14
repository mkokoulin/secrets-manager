package handlers

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/internal/auth"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user models.User) error
	AuthUser(ctx context.Context, user models.User) (*auth.TokenDetails, error)
	DeleteUser(ctx context.Context, userID string) error
	RefreshToken(ctx context.Context, refreshToken string) (*auth.TokenDetails, error)
}

type SecretServiceInterface interface {
	AddSecret(ctx context.Context, secret models.Secret) error
	GetSecrets(ctx context.Context, userID string) ([]models.SecretData, error)
	GetSecret(ctx context.Context, userID string) (models.SecretData, error)
	UpdateSecret(ctx context.Context, secretID, userID string) (models.SecretData, error)
	DeleteSecret(ctx context.Context, secretID, userID string) error
}