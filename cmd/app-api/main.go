package main

import (
	"context"
	"flag"

	"github.com/huseinnashr/bimble/internal/config"
	"github.com/huseinnashr/bimble/internal/pkg/tracer"
)

const APP_NAME = "bimble-backend-http"

func main() {
	var configPath string
	var ctx = context.Background()

	flag.StringVar(&configPath, "config", "./files/config/local.yaml", "path to config file")
	flag.Parse()

	config, err := config.GetConfig(configPath)
	if err != nil {
		panic(err)
	}

	tracerShutdown, err := tracer.Init(ctx, config, APP_NAME)
	if err != nil {
		panic(err)
	}
	defer tracerShutdown(ctx)

	if err := startApp(ctx, config); err != nil {
		panic(err)
	}
}
