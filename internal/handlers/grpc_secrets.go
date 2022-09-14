package handlers

import (
	"context"

	pb "github.com/mkokoulin/secrets-manager.git/internal/pb/secrets"
)

type GRPCSecrets struct {
	pb.UnimplementedUsersServer
	secretService SecretServiceInterface
}

func NewGRPCSecrets(secretService SecretServiceInterface) *GRPCSecrets {
	return &GRPCSecrets {
		secretService: secretService,
	}
}

func (gs *GRPCSecrets) CreateSecret(ctx context.Context, in *pb.CreateSecretRequest) (*pb.CreateSecretResponse, error) {
	gs.secretService.AddSecret(ctx, in.Data)
}