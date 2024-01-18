package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/huseinnashr/bimble/api/v1"
	"github.com/huseinnashr/bimble/internal/config"
	"golang.org/x/sync/errgroup"
)

type IServer interface {
	Start(context.Context) error
	Stop(context.Context) error
}

func startServer(ctx context.Context, config *config.Config, accountHandler v1.AccountServiceServer) error {
	var servers []IServer

	httpServer := http.NewServer(
		http.Address(config.Server.HTTP.Address),
	)
	v1.RegisterAccountServiceHTTPServer(httpServer, accountHandler)
	servers = append(servers, httpServer)

	var opts []grpc.ServerOption
	if address := config.Server.GRPC.Address; address != "" {
		opts = append(opts, grpc.Address(address))
	}
	if timeout := config.Server.GRPC.Timeout; timeout != 0 {
		opts = append(opts, grpc.Timeout(timeout))
	}
	grpcServer := grpc.NewServer(opts...)
	v1.RegisterAccountServiceServer(grpcServer, accountHandler)
	servers = append(servers, grpcServer)

	group, ctx := errgroup.WithContext(ctx)
	for _, server := range servers {
		server := server
		group.Go(func() error {
			if err := server.Start(ctx); err != nil {
				return err
			}
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}
