package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huseinnashr/bimble/internal/config"
	accounthandler "github.com/huseinnashr/bimble/internal/handler/api/account"
	accountrepo "github.com/huseinnashr/bimble/internal/repo/account"
	accountusecase "github.com/huseinnashr/bimble/internal/usecase/account"
	_ "github.com/lib/pq"
	redisv9 "github.com/redis/go-redis/v9"
)

func startApp(ctx context.Context, config *config.Config) error {
	sqlDSN := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Resource.SQLDatabase.Host, config.Resource.SQLDatabase.Port, config.Resource.SQLDatabase.User,
		config.Resource.SQLDatabase.Password, config.Resource.SQLDatabase.DBName,
	)
	sqlDatabase, err := sql.Open("postgres", sqlDSN)
	if err != nil {
		return err
	}

	redis := redisv9.NewClient(&redisv9.Options{
		Addr: config.Resource.Redis.Address, Password: config.Resource.Redis.Password,
	})

	accountRepo := accountrepo.New(config, sqlDatabase, redis)
	accountUsecase := accountusecase.New(accountRepo)
	accountHandler := accounthandler.New(accountUsecase)

	if err := startServer(ctx, config, accountHandler); err != nil {
		return err
	}

	return nil
}
