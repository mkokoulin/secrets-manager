// Package middleware includes functions for additional request processing
package middleware

import (
	"context"

	"github.com/mkokoulin/secrets-manager.git/server/internal/auth"
)

type JWTMiddleware struct {
	AccessTokenSecret string
}

func NewJWTMiddleware(accessTokenSecret string) *JWTMiddleware {
	return &JWTMiddleware{
		AccessTokenSecret: accessTokenSecret,
	}
}

// CheckAuth method for checking auth
func (jwt *JWTMiddleware) CheckAuth(ctx context.Context) (context.Context, error) {
	userID, err := auth.TokenValid(ctx, jwt.AccessTokenSecret)
	if err != nil {
		userID = ""
	}
	newCtx := context.WithValue(ctx, auth.ContextValue, userID)
	return newCtx, nil
}
