package services

import (
	"context"
	"encoding/json"

	"github.com/mkokoulin/secrets-manager.git/server/internal/models"
)

// SecretsService structure of the secrets service
type SecretsService struct {
	db SecretsRepoInterface
}

// NewSecretsService function of creating a service for working with secrets
func NewSecretsService(db SecretsRepoInterface) *SecretsService {
	return &SecretsService {
		db: db,
	}
}

// AddSecret method for creating a new secret
func (ss *SecretsService) AddSecret(ctx context.Context, secret models.RawSecretData) error {
	return ss.db.AddSecret(ctx, secret)
}

// GetSecret method of getting a secret of the user
func (ss *SecretsService) GetSecret(ctx context.Context, secretID, userID string) (models.Secret, error) {
	rawSecret, err := ss.db.GetSecret(ctx, secretID, userID)
	if err != nil {
		return models.Secret{}, err
	}

	var secret models.Secret

	err = rawSecret.Decrypt()
	if err != nil {
		return models.Secret{}, err
	}

	var value map[string]string

	err = json.Unmarshal(rawSecret.Data, &value)
	if err != nil {
		return models.Secret{}, err
	}

	secret.Data.Value = value
	secret.Data.CreatedAt = rawSecret.CreatedAt
		secret.Data.Type = rawSecret.Type

	secret.SecretID = rawSecret.ID
	secret.UserID = rawSecret.UserID

	return secret, nil
}

// DeleteSecret method of deleting a secret of the user
func (ss *SecretsService) DeleteSecret(ctx context.Context, secretID, userID string) error {
	err := ss.db.DeleteSecret(ctx, secretID, userID)
	if err != nil {
		return err
	}

	return nil
}