// Package config incudes server configuration
package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

const (
	DefaultAddress = "localhost:5001"
)

type Config struct {
	Address string `env:"ADDRESS"`
}

func New() *Config {
	cfg := Config{
		Address: DefaultAddress,
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.Address, "d", DefaultAddress, "address of gGRPC server")

	flag.Parse()

	return &cfg
}
