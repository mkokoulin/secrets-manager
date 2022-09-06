package services

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/internal/models"
)

type UserRepoInterface interface {
	CreateUser(ctx context.Context, user models.User) error
	CheckUserPassword(ctx context.Context, user models.User) (string, error)
	DeleteUser(ctx context.Context, userID string) error
}