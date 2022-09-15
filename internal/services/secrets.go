package services

import (
	"context"
	"encoding"

	"github.com/mkokoulin/secrets-manager.git/internal/helpers/encryptor"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
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

	var result models.Secret

	encryptor.Decrypt(rawSecret.Data)



	return 
}

func (ss *SecretsService) UpdateSecret(ctx context.Context, secretID, userID string, secret models.SecretData) error {
	return nil
}

func (ss *SecretsService) DeleteSecret(ctx context.Context, secretID, userID string) error {
	return nil
}