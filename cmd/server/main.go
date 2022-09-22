package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/mkokoulin/secrets-manager.git/internal/middleware"
	"github.com/mkokoulin/secrets-manager.git/internal/config"
	"github.com/mkokoulin/secrets-manager.git/internal/database"
	"github.com/mkokoulin/secrets-manager.git/internal/handlers"
	"github.com/mkokoulin/secrets-manager.git/internal/models"
	"github.com/mkokoulin/secrets-manager.git/internal/services"
)

var (
	grpcServer   *grpc.Server
)

func init() {
	zerolog.TimeFieldFormat = "2006.02.01-15:04:05.000"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()

	g, ctx := errgroup.WithContext(ctx)

	interrupt := make(chan os.Signal, 1)

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
	conn.Table("secrets").AutoMigrate(&models.RawSecretData{})

	repo := database.NewPostgresDatabase(conn)

	userService := services.NewUsersService(repo, cfg.AccessTokenLiveTimeMinutes, cfg.RefreshTokenLiveTimeDays, cfg.AccessTokenSecret, cfg.RefreshTokenSecret)
	GRPCUsers := handlers.NewGRPCUsers(userService)
	
	secretsService := services.NewSecretsService(repo)
	GRPCSecrets := handlers.NewGRPCSecrets(secretsService)

	JWTMiddleware := middleware.NewJWTMiddleware(cfg.AccessTokenSecret)

	GRPCServer := services.NewGrpcServer(
		services.WithServerConfig(cfg),
		services.WithServices(GRPCUsers, GRPCSecrets),
		services.WithStreamInterceptors(
			grpcauth.StreamServerInterceptor(JWTMiddleware.CheckAuth),
		),
		services.WithUnaryInterceptors(
			grpcauth.UnaryServerInterceptor(JWTMiddleware.CheckAuth),
		),
	)

	g.Go(func() error {
		err := GRPCServer.Start(cancel)
		if err != nil {
			log.Error().Caller().Str("gRPC server failed to listen", "").Err(err).Msg("")
			return err
		}
		
		return nil
	})

	select {
	case <-interrupt:
		log.Debug().Msgf("stop server")
		break
	case <-ctx.Done():
		// GRPCServer.Stop()

		log.Log().Caller().Msg("Stop server")
		break
	}

	err = g.Wait()
	if err != nil {
		log.Error().Caller().Str("server returning an error: ", err.Error()).Err(err).Msg("")
	}


	log.Log().Caller().Msg("Run server")
}
