package services

import (
	"context"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/mkokoulin/secrets-manager.git/internal/config"
)

type Service interface {
	RegisterService(grpc.ServiceRegistrar)
}

type GRPCServer struct {
	cfg config.Config
	server *grpc.Server
	logger             *zerolog.Logger
	services           []Service
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
}

type GrpcServerOption func(*GRPCServer)

func NewGrpcServer(opts ...GrpcServerOption) *GRPCServer {
	s := &GRPCServer{}

	for _, option := range opts {
		option(s)
	}

	return s
}

func WithStreamInterceptors(in ...grpc.StreamServerInterceptor) GrpcServerOption {
	return func(server *GRPCServer) {
		server.streamInterceptors = append(server.streamInterceptors, in...)
	}
}

func WithUnaryInterceptors(in ...grpc.UnaryServerInterceptor) GrpcServerOption {
	return func(server *GRPCServer) {
		server.unaryInterceptors = append(server.unaryInterceptors, in...)
	}
}

func WithServices(s ...Service) GrpcServerOption {
	return func(server *GRPCServer) {
		server.services = s
	}
}

func (s *GRPCServer) RegisterServices(services ...Service) {
	for _, service := range services {
		service.RegisterService(s.server)
	}
}

func (s *GRPCServer) Start(cancel context.CancelFunc) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GRPCPort))
	if err != nil {
		s.logger.Error().Caller().Str("gRPC server failed to listen", "").Err(err).Msg("")
		return err
	}

	s.server = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			s.streamInterceptors...,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			s.unaryInterceptors...,
		)),
	)

	s.RegisterServices(s.services...)

	err = s.server.Serve(lis)
	if err != nil {
		s.logger.Error().Caller().Str("gRPC server failed to listen", "").Err(err).Msg("")
		cancel()
	}

	s.logger.Log().Caller().Msgf("gRPC server is running on %s port", s.cfg.GRPCPort)

	return nil
}

func (s *GRPCServer) Stop() {
	s.logger.Log().Caller().Msg("Gracefully stopping gRPC server")

	s.server.GracefulStop()

	s.logger.Log().Caller().Msg("gRPC server stopped")
}