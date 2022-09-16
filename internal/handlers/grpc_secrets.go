package handlers

import (
	"context"

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
	data := map[string]string {}

	for _, v := range in.Data {
		data[v.Title] = v.Value
	}

	secret := models.Secret{
		UserID: "",
		Data: models.SecretData {
			Type: in.Type,
			Value: data,
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
	secret, err := gs.secretService.GetSecret(ctx, in.SecretId, "")
	if err != nil {
		return nil, customerrors.NewCustomError(err, "")
	}

	dates := []*pb.Data {}

	for k, v := range secret.Data.Value {
		data := pb.Data{}
		
		data.Title = k
		data.Value = v

		dates = append(dates, &data)
	}

	return &pb.GetSecretResponse {
		Id: secret.SecretID,
		Type: secret.Data.Type,
		Status: "ok",
		Secret: dates,
	}, nil
}