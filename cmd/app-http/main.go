package main

import (
	"context"
	"flag"

	"github.com/huseinnashr/bimble/internal/config"
)

func main() {
	var configPath string
	var ctx = context.Background()

	flag.StringVar(&configPath, "config", "./files/config/local.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.GetConfig(configPath)
	if err != nil {
		panic(err)
	}

	if err := startApp(ctx, cfg); err != nil {
		panic(err)
	}
}
