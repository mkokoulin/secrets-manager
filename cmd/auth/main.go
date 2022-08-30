package main

import (
	"flag"
	"net"
	url2 "net/url"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var (
	socket   = flag.String("socket", "tcp://0.0.0.0:8080", "auth-service address to run")
	logLevel = flag.String("log-level", "info", "logging level")
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func main() {
	flag.Parse()

	l, err := zerolog.ParseLevel(*logLevel)
	if err != nil {
		log.Error().Caller().Str("log-level", *logLevel).Err(err).Msg("")
		return
	}
	zerolog.SetGlobalLevel(l)

	s := grpc.NewServer()

	url, err := url2.Parse(*socket)
	if err != nil {
		log.Error().Caller().Str("socket", *socket).Err(err).Msg("")
		return
	}

	protocol := url.Scheme
	address := url.Host
}
