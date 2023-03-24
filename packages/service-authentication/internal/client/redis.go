package client

import (
	"context"
	"os"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/redis/go-redis/v9"
)

func getRedisClient() (*redis.Client, *customerror.CustomError) {
	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")
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
