package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mkokoulin/secrets-manager.git/internal/config"
	"github.com/mkokoulin/secrets-manager.git/internal/database"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
)

func init() {
	zerolog.TimeFieldFormat = "2006.02.01-15:04:05.000"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func main() {
	cfg := config.New()

	l, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Error().Caller().Str("log-level", cfg.LogLevel).Err(err).Msg("")
		return
	}
	zerolog.SetGlobalLevel(l)

	log.Log().Msg("=========================")
	log.Log().Msgf("Database URI: %s", cfg.DatabaseURI)
	log.Log().Msgf("gRPC port:    %d", cfg.GRPCPort)
	log.Log().Msgf("Log level:    %s", cfg.LogLevel)
	log.Log().Msg("=========================")

	conn, err := gorm.Open(postgres.Open(cfg.DatabaseURI), &gorm.Config{})
	if err != nil {
		log.Error().Caller().Str("database URI", cfg.DatabaseURI).Err(err).Msg("")
	}

	conn.AutoMigrate(&models.User{})

	p := database.NewPostgresDatabase(conn)

	userID, err := p.CreateUser(context.Background(), models.User{
		ID: uuid.New(),
		Login: "string1",
		Password: "string2",
		CreatedAt: time.Now(),
		IsDeleted: false,
	})
	if err != nil {
		log.Error().Caller().Str("create user", "").Err(err).Msg("")
	}

	err = p.DeleteUser(context.Background(), userID)
	if err != nil {
		log.Error().Caller().Str("delete user", "").Err(err).Msg("")
	}


	log.Log().Caller().Msg("Run server")

	// token, err := auth.CreateToken("UserID", cfg.AccessTokenSecret, cfg.RefreshTokenSecret, cfg.AccessTokenLiveTimeMinutes, cfg.RefreshTokenLiveTimeDays)
	// if err != nil {
	// 	return
	// }

	// fmt.Println(token)
}
