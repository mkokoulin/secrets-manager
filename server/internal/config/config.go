package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

const (
	DefaultDatabaseURI = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	DefaultGRPCPort = 5001
	DefaultLogLevel = "info"
	DefaultAccessTokenSecret = "fgdghjafdgdfg"
	DefaultRefreshTokenSecret = "ofdfgdgjhfghfghjghjsw"
	DefaultAccessTokenLiveTimeMinutes = 10
	DefaultRefreshTokenLiveTimeDays = 1
)

type Config struct {
	DatabaseURI string `env:"DATABASE_URI"`
	GRPCPort int `env:"GRPC_PORT"`
	LogLevel string `env:"LOG_LEVEL"`	
	AccessTokenSecret string `env:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `env:"REFRESH_TOKEN_SECRET"`
	AccessTokenLiveTimeMinutes int `env:"ACCESS_TOKEN_LIVE_TIME_MINUTES"`
	RefreshTokenLiveTimeDays int `env:"REFRESH_TOKEN_LIVE_TIME_DAYS"`
}

func New() *Config {
	cfg := Config{
		DatabaseURI: DefaultDatabaseURI,
		GRPCPort: DefaultGRPCPort,
		LogLevel: DefaultLogLevel,
		AccessTokenSecret: DefaultAccessTokenSecret,
		RefreshTokenSecret: DefaultRefreshTokenSecret,
		AccessTokenLiveTimeMinutes: DefaultAccessTokenLiveTimeMinutes,
		RefreshTokenLiveTimeDays: DefaultRefreshTokenLiveTimeDays,
	}


	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.DatabaseURI, "d", DefaultDatabaseURI, "database URI")

	flag.IntVar(&cfg.GRPCPort, "g", DefaultGRPCPort, "gRPC port")

	flag.StringVar(&cfg.LogLevel, "l", DefaultLogLevel, "logging level")

	flag.Parse()

	return &cfg
}