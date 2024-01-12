package main

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/huseinnashr/bimble/api/v1"
	"github.com/huseinnashr/bimble/internal/config"
)

func startServer(ctx context.Context, config *config.Config, accountHandler v1.AccountServiceHTTPServer) error {
	server := http.NewServer(
		http.Address(fmt.Sprintf(":%d", config.Port)),
	)

	v1.RegisterAccountServiceHTTPServer(server, accountHandler)

	if err := server.Start(ctx); err != nil {
		return err
	}

	return nil
}
