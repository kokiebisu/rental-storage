package client

import (
	"context"
	"os"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/redis/go-redis/v9"
)

func getRedisClient() (*redis.Client, *customerror.CustomError) {
	addr := os.Getenv("ELASTICACHE_HOST")
	port := os.Getenv("ELASTICACHE_PORT")
	endpoint := addr + ":" + port
	db := 0

	client := redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: "",
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, customerror.ErrorHandler.RedisConnectionError(err)
	}

	return client, nil
}

func getRedisDevelopmentClient() (*redis.Client, *customerror.CustomError) {
	addr := "localhost:6379"
	password := ""
	db := 0

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, customerror.ErrorHandler.RedisConnectionError(err)
	}

	return client, nil
}
