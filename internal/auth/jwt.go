package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
		RtExpires: time.Now().Add(time.Second - time.Duration(refreshTokenLiveTimeDays)).Unix(),
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

func ValidateToken(r *http.Request, accessTokenSecret string) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New("expired token")
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
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
