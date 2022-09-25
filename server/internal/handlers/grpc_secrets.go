// Package handlers wrappers for working with pb files
package handlers

import (
	"context"

	customerrors "github.com/mkokoulin/secrets-manager.git/server/internal/errors"
	"github.com/mkokoulin/secrets-manager.git/server/internal/models"
	pb "github.com/mkokoulin/secrets-manager.git/server/internal/pb/secrets"
	"google.golang.org/grpc"
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

func (gs *GRPCSecrets) RegisterService(r grpc.ServiceRegistrar) {
	pb.RegisterSecretsServer(r, gs)
}

func (gs *GRPCSecrets) CreateSecret(ctx context.Context, in *pb.CreateSecretRequest) (*pb.CreateSecretResponse, error) {
	userID := getUserFromContext(ctx)
	if userID == "" {
		return &pb.CreateSecretResponse{
			Status: "unauthorized",
		}, nil
	}

	data := map[string]string {}

	for _, v := range in.Data {
		data[v.Title] = v.Value
	}

	secret := models.Secret{
		UserID: userID,
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
		return nil, customerrors.NewCustomError(err, "an error occurred while saving the secret")
	}

	return &pb.CreateSecretResponse {
		Status: "ok",
	}, nil
}

func (gs *GRPCSecrets) GetSecret(ctx context.Context, in *pb.GetSecretRequest)(*pb.GetSecretResponse, error) {
	userID := getUserFromContext(ctx)
	if userID == "" {
		return &pb.GetSecretResponse{
			Status: "unauthorized",
		}, nil
	}

	secret, err := gs.secretService.GetSecret(ctx, in.SecretId, userID)
	if err != nil {
		return nil, customerrors.NewCustomError(err, "an error occurred while receiving the secret")
	}

	date := []*pb.Data {}

	for k, v := range secret.Data.Value {
		d := pb.Data{}
		
		d.Title = k
		d.Value = v

		date = append(date, &d)
	}

	return &pb.GetSecretResponse {
		Id: secret.SecretID,
		Type: secret.Data.Type,
		Status: "ok",
		Secret: date,
	}, nil
}

func (gs *GRPCSecrets) DeleteSecret(ctx context.Context, in *pb.DeleteSecretRequest)(*pb.DeleteSecretResponse, error) {
	userID := getUserFromContext(ctx)
	if userID == "" {
		return &pb.DeleteSecretResponse{
			Status: "unauthorized",
		}, nil
	}

	err := gs.secretService.DeleteSecret(ctx, in.SecretId, userID)
	if err != nil {
		return nil, customerrors.NewCustomError(err, "an error occurred while deleting the secret")
	}

	return &pb.DeleteSecretResponse {
		Status: "ok",
	}, nil
}

func getUserFromContext(ctx context.Context) string {
	userID := ctx.Value("userID")
	if userID != nil {
		if str, ok := userID.(string); ok {
			return str
		} else {
			return ""
		}
	}
	return ""
}