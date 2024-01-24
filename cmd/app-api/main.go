package main

import (
	"context"
	"flag"

	"github.com/huseinnashr/bimble/internal/config"
	"github.com/huseinnashr/bimble/internal/pkg/tracer"
)

var (
	Name    string = "bimble-backend-api"
	Version string
)

func main() {
	var configPath string
	var ctx = context.Background()

	flag.StringVar(&configPath, "config", "./files/config/local.yaml", "path to config file")
	flag.Parse()

	config, err := config.GetConfig(configPath)
	if err != nil {
		panic(err)
	}
	config.App.Name = Name
	config.App.Version = Version

	tracerShutdown, err := tracer.Init(ctx, config)
	if err != nil {
		panic(err)
	}
	defer tracerShutdown(ctx)

	if err := startApp(ctx, config); err != nil {
		panic(err)
	}
}
