package services

import (
	"context"

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

func (ss *SecretsService) AddSecret(ctx context.Context, secret models.Secret) error {
	return ss.db.AddSecret(ctx, secret)
}