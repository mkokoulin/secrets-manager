package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/mkokoulin/secrets-manager.git/client/internal/handlers"
	"github.com/mkokoulin/secrets-manager.git/client/internal/loops"
	"github.com/mkokoulin/secrets-manager.git/client/internal/services"
)

var (
	BuildTime  string
	AppVersion string
	address    string
)

const defaultAddress = "localhost:50051"

func init() {
	envAddress := os.Getenv("ADDRESS")
	if envAddress != "" {
		address = envAddress
	}
	address = *flag.String("a", defaultAddress, "address of gGRPC server")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("InitApp")
	fmt.Printf("App version: %v, Date compile: %v\n", AppVersion, BuildTime)
	userClient := services.NewUserClient(address)
	userHandler := handlers.NewUserHandler(userClient)

	userLoop := loops.NewUserLoop(address, userHandler)
	userLoop.MainLoop(ctx)
	cancel()
}