// Package main entrance file for client application
package main

import (
	"context"
	"fmt"

	"github.com/mkokoulin/secrets-manager.git/client/internal/handlers"
	"github.com/mkokoulin/secrets-manager.git/client/internal/loops"
	"github.com/mkokoulin/secrets-manager.git/client/internal/services"
	"github.com/mkokoulin/secrets-manager.git/client/internal/config"
)

var (
	BuildTime  string
	AppVersion string
)

func main() {
	cfg := config.New()

	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("InitApp")
	fmt.Printf("App version: %v, Date compile: %v\n", AppVersion, BuildTime)
	userClient := services.NewUserClient(cfg.Address)
	userHandler := handlers.NewUserHandler(userClient)

	userLoop := loops.NewUserLoop(cfg.Address, userHandler)
	userLoop.MainLoop(ctx)
	cancel()
}