package services

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/internal/auth"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
)

type UsersService struct {
	db UserRepoInterface
	AccessTokenLiveTimeMinutes int
	RefreshTokenLiveTimeDays   int
	AccessTokenSecret          string
	RefreshTokenSecret         string
}

func NewUsersService(db UserRepoInterface, accessTokenLiveTimeMinutes, refreshTokenLiveTimeDays int, accessTokenSecret, refreshTokenSecret string) *AuthService {
	return &UsersService{
		db: db,
		AccessTokenLiveTimeMinutes: accessTokenLiveTimeMinutes,
		RefreshTokenLiveTimeDays: refreshTokenLiveTimeDays,
		AccessTokenSecret: accessTokenSecret,
		RefreshTokenSecret: refreshTokenSecret,
	}
}

func (us *UsersService) CreateUser(ctx context.Context, user models.User) error {
	return us.db.CreateUser(ctx, user)
}

func (us *UsersService) LoginUser(ctx context.Context, user models.User)(*auth.TokenDetails, error) {
	userID, err := us.db.CheckUserPassword(ctx, user)
	if err != nil {
		return nil, err
	}

	return auth.CreateToken(userID, us.AccessTokenSecret, us.RefreshTokenSecret, us.AccessTokenLiveTimeMinutes, us.RefreshTokenLiveTimeDays)
}

func (us *UsersService) DeleteUser(ctx context.Context, userID string) error {
	return us.db.DeleteUser(ctx, userID)
}

func (us *UsersService) RefreshToken(ctx context.Context, refreshToken string) (*auth.TokenDetails, error) {
	return auth.RefreshToken(refreshToken, us.RefreshTokenSecret, us.AccessTokenLiveTimeMinutes, us.RefreshTokenLiveTimeDays)
}