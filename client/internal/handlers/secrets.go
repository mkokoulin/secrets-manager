// Package handlers includes methods for handling users cases
package handlers

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/client/internal/models"
	"github.com/mkokoulin/secrets-manager.git/client/internal/services"
)

// NewSecretHandler function for creates new secret handler
func NewSecretHandler(client *services.SecretClient) *SecretHandler {
	return &SecretHandler{
		secretClient: client,
	}
}

// SecretHandler structure for handling secret
type SecretHandler struct {
	secretClient *services.SecretClient
	userClient   *services.UserClient
}

// CreateSecret method of creating a secret
func (sh *SecretHandler) CreateSecret(ctx context.Context, secret models.Secret) error {
	return sh.secretClient.CreateSecret(ctx, secret)
}

// GetSecret method of getting a secret
func (sh *SecretHandler) GetSecret(ctx context.Context) (models.Secret, error) {
	return sh.secretClient.GetSecret(ctx)
}