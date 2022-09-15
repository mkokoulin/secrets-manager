package handlers

import (
	"context"
	"encoding/json"

	customerrors "github.com/mkokoulin/secrets-manager.git/internal/errors"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
	pb "github.com/mkokoulin/secrets-manager.git/internal/pb/secrets"
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
		Data: models.SecretData {
			Type: in.Type,
			Value: value,
		},
	}

	rawSecret, err := models.NewRawSecretData(secret)
	if err != nil {
		return nil, err
	}

	err = gs.secretService.AddSecret(ctx, *rawSecret)
	if err != nil {
		return nil, customerrors.NewCustomError(err, "")
	}

	return &pb.CreateSecretResponse {
		Status: "ok",
	}, nil
}

func (gs *GRPCSecrets) GetSecret(ctx context.Context, in *pb.GetSecretRequest)(*pb.GetSecretResponse, error) {
	rawSecret, err := gs.secretService.GetSecret(ctx, in.SecretId, "")
	if err != nil {
		return nil, customerrors.NewCustomError(err, "")
	}

	var d map[string]string

	err = json.Unmarshal(rawSecret.Data, &d)
	if err != nil {
		return nil, customerrors.NewCustomError(err, "")
	}

	return &pb.GetSecretResponse {
		Id: rawSecret.ID,
		Type: rawSecret.Type,
		Data: &pb.Data{
			Title: ,
			Value: "",
		},
		Status: "ok",
	}, nil
}