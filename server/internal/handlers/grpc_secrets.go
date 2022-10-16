// Package handlers wrappers for working with pb files
package handlers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/mkokoulin/secrets-manager.git/server/internal/services"
	"github.com/mkokoulin/secrets-manager.git/server/internal/auth"
	"github.com/mkokoulin/secrets-manager.git/server/internal/models"

	customerrors "github.com/mkokoulin/secrets-manager.git/server/internal/errors"
	pb "github.com/mkokoulin/secrets-manager.git/server/internal/pb/secrets"
)

type GRPCSecrets struct {
	pb.UnimplementedSecretsServer
	secretService SecretServiceInterface
}

func NewGRPCSecrets(secretService SecretServiceInterface) *GRPCSecrets {
	return &GRPCSecrets{
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
			Status: services.UnauthorizedStatus,
		}, nil
	}

	data := map[string]string{}

	for _, v := range in.Data {
		data[v.Title] = v.Value
	}

	secret := models.Secret{
		UserID: userID,
		Data: models.SecretData{
			Type:  in.Type,
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

	return &pb.CreateSecretResponse{
		Status: "ok",
	}, nil
}

func (gs *GRPCSecrets) GetSecret(ctx context.Context, in *pb.GetSecretRequest) (*pb.GetSecretResponse, error) {
	userID := getUserFromContext(ctx)
	if userID == "" {
		return &pb.GetSecretResponse{
			Status: services.UnauthorizedStatus,
		}, nil
	}

	secret, err := gs.secretService.GetSecret(ctx, in.SecretId, userID)
	if err != nil {
		return nil, customerrors.NewCustomError(err, "an error occurred while receiving the secret")
	}

	date := []*pb.Data{}

	for k, v := range secret.Data.Value {
		d := pb.Data{}

		d.Title = k
		d.Value = v

		date = append(date, &d)
	}

	return &pb.GetSecretResponse{
		Id:     secret.SecretID,
		Type:   secret.Data.Type,
		Status: "ok",
		Secret: date,
	}, nil
}

func (gs *GRPCSecrets) GetSecrets(ctx context.Context, in *pb.GetSecretsRequest) (*pb.GetSecretsResponse, error) {
	userID := getUserFromContext(ctx)
	if userID == "" {
		return &pb.GetSecretsResponse{
			Status: services.UnauthorizedStatus,
		}, nil
	}

	secrets, err := gs.secretService.GetSecrets(ctx, userID)
	if err != nil {
		return &pb.GetSecretsResponse{
			Status: "error",
		}, nil
	}

	s := []*pb.GetSecretsResponse_Secret{}

	for _, v := range secrets {
		var secretResponse pb.GetSecretsResponse_Secret
		secretResponseData := []*pb.Data{}

		secretResponse.Id = v.SecretID

		for k, v := range v.Data.Value {
			d := pb.Data{}

			d.Title = k
			d.Value = v

			secretResponseData = append(secretResponseData, &d)
		}

		secretResponse.Data = secretResponseData

		s = append(s, &secretResponse)
	}

	return &pb.GetSecretsResponse{
		Status:  "ok",
		Secrets: s,
	}, nil
}

func (gs *GRPCSecrets) DeleteSecret(ctx context.Context, in *pb.DeleteSecretRequest) (*pb.DeleteSecretResponse, error) {
	userID := getUserFromContext(ctx)
	if userID == "" {
		return &pb.DeleteSecretResponse{
			Status: services.UnauthorizedStatus,
		}, nil
	}

	err := gs.secretService.DeleteSecret(ctx, in.SecretId, userID)
	if err != nil {
		return nil, customerrors.NewCustomError(err, "an error occurred while deleting the secret")
	}

	return &pb.DeleteSecretResponse{
		Status: "ok",
	}, nil
}

func getUserFromContext(ctx context.Context) string {
	userID := ctx.Value(auth.ContextValue)
	if userID != nil {
		if str, ok := userID.(string); ok {
			return str
		}
	}
	return ""
}
