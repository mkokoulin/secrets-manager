package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
 	"strings"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/golang-jwt/jwt/v4"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(userID string, accessTokenSecret, refreshTokenSecret string, accessTokenLiveTimeMinutes, refreshTokenLiveTimeDays int) (*TokenDetails, error) {
	td := &TokenDetails{
		AtExpires: time.Now().Add(time.Minute * time.Duration(accessTokenLiveTimeMinutes)).Unix(),
		RtExpires: time.Now().Add(time.Hour * 24 * time.Duration(refreshTokenLiveTimeDays)).Unix(),
	}

	atClaims := jwt.MapClaims{
		"exp":     td.AtExpires,
		"user_id": userID,
	}

	rtClaims := jwt.MapClaims{
		"exp":     td.RtExpires,
		"user_id": userID,
	}

	atWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	rtWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	at, err := atWithClaims.SignedString([]byte(accessTokenSecret))
	if err != nil {
		return nil, err
	}

	rt, err := rtWithClaims.SignedString([]byte(refreshTokenSecret))
	if err != nil {
		return nil, err
	}

	td.AccessToken = at
	td.RefreshToken = rt

	log.Println("token has been generated")

	return td, nil
}

func ValidateToken(ctx context.Context, accessTokenSecret string) (*jwt.Token, error) {
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(accessTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(ctx context.Context, accessSecret string) (string, error) {
	token, err := ValidateToken(ctx, accessSecret)
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", err
	}
	mapClaims := token.Claims.(jwt.MapClaims)
	t := mapClaims["user_id"].(string)
	return t, nil
}

func ExtractToken(ctx context.Context) string {
	token := metautils.ExtractIncoming(ctx).Get("authorization")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func RefreshToken(refreshToken, refreshTokenSecret string, accessTokenLiveTimeMinutes, refreshTokenLiveTimeDays int) (*TokenDetails, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userID := claims["user_id"].(string)

		td, err := CreateToken(userID, refreshToken, refreshTokenSecret, accessTokenLiveTimeMinutes, refreshTokenLiveTimeDays)
		if err != nil {
			return nil, err
		}

		log.Println("token has been refreshed")

		return td, nil
	} else {
		return nil, errors.New("refresh token expired")
	}
}
