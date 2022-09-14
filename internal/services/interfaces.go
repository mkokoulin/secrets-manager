package services

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/internal/models"
)

type UserRepoInterface interface {
	CreateUser(ctx context.Context, user models.User) error
	CheckUserPassword(ctx context.Context, user models.User) (string, error)
	DeleteUser(ctx context.Context, userID string) error
}

type SecretsRepoInterface interface {
	AddSecret(ctx context.Context, secret models.Secret) error
	GetSecret(ctx context.Context, secretID, userID string) (models.SecretData, error)
	GetSecrets(ctx context.Context, userID string) ([]models.SecretData, error)
	UpdateSecret(ctx context.Context, secretID, userID string, secret models.SecretData) error
	DeleteSecret(ctx context.Context, secretID, userID string) error
}