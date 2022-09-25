package services

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/server/internal/models"
)

// UserRepoInterface interface for interacting with the user database
type UserRepoInterface interface {
	CreateUser(ctx context.Context, user models.User) error
	CheckUserPassword(ctx context.Context, user models.User) (string, error)
	DeleteUser(ctx context.Context, userID string) error
}

// SecretsRepoInterface interface for interacting with a database of secrets
type SecretsRepoInterface interface {
	AddSecret(ctx context.Context, secret models.RawSecretData) error
	GetSecret(ctx context.Context, secretID, userID string) (models.RawSecretData, error)
	GetSecrets(ctx context.Context, userID string) ([]models.SecretData, error)
	UpdateSecret(ctx context.Context, secretID, userID string, secret models.SecretData) error
	DeleteSecret(ctx context.Context, secretID, userID string) error
}