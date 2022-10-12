package services

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/server/internal/auth"
	"github.com/mkokoulin/secrets-manager.git/server/internal/models"
)

// UsersService structure for the user service
type UsersService struct {
	db UserRepoInterface
	AccessTokenLiveTimeMinutes int
	RefreshTokenLiveTimeDays   int
	AccessTokenSecret          string
	RefreshTokenSecret         string
}

// NewUsersService function of creating a service for working with users
func NewUsersService(db UserRepoInterface, accessTokenLiveTimeMinutes, refreshTokenLiveTimeDays int, accessTokenSecret, refreshTokenSecret string) *UsersService {
	return &UsersService{
		db: db,
		AccessTokenLiveTimeMinutes: accessTokenLiveTimeMinutes,
		RefreshTokenLiveTimeDays: refreshTokenLiveTimeDays,
		AccessTokenSecret: accessTokenSecret,
		RefreshTokenSecret: refreshTokenSecret,
	}
}

// CreateUser user creation method
func (us *UsersService) CreateUser(ctx context.Context, user models.User) error {
	return us.db.CreateUser(ctx, user)
}

// AuthUser user authorization method
func (us *UsersService) AuthUser(ctx context.Context, user models.User)(*auth.TokenDetails, error) {
	userID, err := us.db.CheckUserPassword(ctx, user)
	if err != nil {
		return nil, err
	}

	return auth.CreateToken(userID, us.AccessTokenSecret, us.RefreshTokenSecret, us.AccessTokenLiveTimeMinutes, us.RefreshTokenLiveTimeDays)
}

// DeleteUser user deletion method
func (us *UsersService) DeleteUser(ctx context.Context, userID string) error {
	return us.db.DeleteUser(ctx, userID)
}

// RefreshToken method for updating user tokens
func (us *UsersService) RefreshToken(refreshToken string) (*auth.TokenDetails, error) {
	return auth.RefreshToken(refreshToken, us.RefreshTokenSecret, us.AccessTokenLiveTimeMinutes, us.RefreshTokenLiveTimeDays)
}