package services

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/mkokoulin/secrets-manager.git/client/internal/models"
	pb "github.com/mkokoulin/secrets-manager.git/client/internal/pb/secrets"
)

type GRPCClient struct {
	pb.SecretsClient
	closeFunc func() error
}

// SecretClient структура клиента для работы с секретами
type SecretClient struct {
	address      string
	accessToken  string
	refreshToken string
	userClient   *UserClient
}

func NewSecretClient(address string, access string, refresh string, userClient *UserClient) *SecretClient {
	return &SecretClient{
		address:      address,
		accessToken:  access,
		refreshToken: refresh,
		userClient:   userClient,
	}
}

func (c *SecretClient) GetSecret(ctx context.Context) (models.Secret, error) {
	client, err := c.getConn()
	if err != nil {
		return models.Secret{}, err
	}

	message := pb.GetSecretRequest{}

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", fmt.Sprintf("Bearer %v", c.accessToken))
	response, err := client.GetSecret(ctx, &message)
	if err != nil {
		return models.Secret{}, err
	}

	var result models.Secret
	if response.Status == "unauthorized" {
		err = c.tryToRefreshToken(ctx)
		if err != nil {
			return models.Secret{}, err
		}
		response, err = client.GetSecret(ctx, &message)
		if err != nil {
			return models.Secret{}, err
		}
	}
	if response.Status != "ok" {
		return models.Secret{}, errors.New(response.Status)
	}

	result.ID = response.Id
	result.Type = response.Type

	for _, secret := range response.Secret {		
		result.Value[secret.Title] = secret.Value
	}

	return result, nil
}

func (c *SecretClient) CreateSecret(ctx context.Context, secret models.Secret) error {
	client, err := c.getConn()
	if err != nil {
		return err
	}
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", fmt.Sprintf("Bearer %v", c.accessToken))
	message := pb.CreateSecretRequest{
		Type:     secret.Type,
 		Data:     secret.TransferValueData(),
	}
	response, err := client.CreateSecret(ctx, &message)
	if err != nil {
		return err
	}
	if response.Status == "unauthorized" {
		err = c.tryToRefreshToken(ctx)
		if err != nil {
			return err
		}
		response, err = client.CreateSecret(ctx, &message)
		if err != nil {
			return err
		}
	}
	if response.Status != "ok" {
		return errors.New(response.Status)
	}
	return nil
}

func (c *SecretClient) tryToRefreshToken(ctx context.Context) error {
	access, refresh, err := c.userClient.Refresh(ctx, c.refreshToken)
	if err != nil {
		return errors.New("failed to refresh token")
	}
	c.accessToken = access
	c.refreshToken = refresh
	return nil
}

func (c *SecretClient) getConn() (*GRPCClient, error) {
	conn, err := grpc.Dial(c.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	cl := pb.NewSecretsClient(conn)

	return &GRPCClient{cl, conn.Close}, nil
}