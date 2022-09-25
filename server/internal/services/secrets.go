package services

import (
	"context"
	"encoding/json"

	"github.com/mkokoulin/secrets-manager.git/server/internal/models"
)

type SecretsService struct {
	db SecretsRepoInterface
}

func NewSecretsService(db SecretsRepoInterface) *SecretsService {
	return &SecretsService {
		db: db,
	}
}

func (ss *SecretsService) AddSecret(ctx context.Context, secret models.RawSecretData) error {
	return ss.db.AddSecret(ctx, secret)
}

func (ss *SecretsService) GetSecrets(ctx context.Context, userID string) ([]models.SecretData, error) {
	return nil, nil
}

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

func (ss *SecretsService) UpdateSecret(ctx context.Context, secretID, userID string, secret models.SecretData) error {
	return nil
}

func (ss *SecretsService) DeleteSecret(ctx context.Context, secretID, userID string) error {
	return nil
}