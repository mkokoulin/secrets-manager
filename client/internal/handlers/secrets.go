package handlers

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/client/internal/models"
	"github.com/mkokoulin/secrets-manager.git/client/internal/services"
)

func NewSecretHandler(client *services.SecretClient) *SecretHandler {
	return &SecretHandler{
		secretClient: client,
	}
}

type SecretHandler struct {
	secretClient *services.SecretClient
	userClient   *services.UserClient
}

func (sh *SecretHandler) CreateSecret(ctx context.Context, secret models.Secret) error {
	return sh.secretClient.CreateSecret(ctx, secret)
}

func (sh *SecretHandler) GetSecret(ctx context.Context) (models.Secret, error) {
	return sh.secretClient.GetSecret(ctx)
}