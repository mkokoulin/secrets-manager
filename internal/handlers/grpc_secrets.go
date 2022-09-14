package handlers

import (
	"context"
	"time"

	"github.com/mkokoulin/secrets-manager.git/internal/models"
	pb "github.com/mkokoulin/secrets-manager.git/internal/pb/secrets"
	customerrors "github.com/mkokoulin/secrets-manager.git/internal/errors"
)

type GRPCSecrets struct {
	pb.UnimplementedSecretsServer
	secretService SecretServiceInterface
}

func NewGRPCSecrets(secretService SecretServiceInterface) *GRPCSecrets {
	return &GRPCSecrets {
		secretService: secretService,
	}
}

func (gs *GRPCSecrets) CreateSecret(ctx context.Context, in *pb.CreateSecretRequest) (*pb.CreateSecretResponse, error) {
	value := map[string]string {
		"title": in.Data.Title,
		"value": in.Data.Value,
	}

	secret := models.Secret{
		UserID: "",
		SecretID: "",
		Data: models.SecretData {
			CreatedAt: time.Now(),
			Type: in.Type,
			Value: value,
		},
	}

	err := gs.secretService.AddSecret(ctx, secret)
	if err != nil {
		return nil, customerrors.NewCustomError(err, "")
	}

	return &pb.CreateSecretResponse {
		Status: "ok",
	}, nil
}