package main

import (
	"context"

	"github.com/huseinnashr/bimble/internal/config"
	accounthttphandler "github.com/huseinnashr/bimble/internal/handler/http/account"
	accountrepo "github.com/huseinnashr/bimble/internal/repo/account"
	accountusecase "github.com/huseinnashr/bimble/internal/usecase/account"
)

func startApp(ctx context.Context, cfg *config.Config) {
	accountRepo := accountrepo.New()
	accountUsecase := accountusecase.New(accountRepo)
	accountHttpHandler := accounthttphandler.New(accountUsecase)

	startServer(ctx, cfg, accountHttpHandler)
}
