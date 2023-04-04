package store

import (
	"context"
	"time"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/redis/go-redis/v9"
)

type TokenStore struct {
	client *redis.Client
}

var (
	AccessTokenGroup  = "access_token"
	RefreshTokenGroup = "refresh_token"
)

// NewTokenStore creates a new token store
func NewTokenStore(client *redis.Client) port.TokenStore {
	return &TokenStore{
		client,
	}
}

// SetAccessToken stores the access token for the user
func (r *TokenStore) SetAccessToken(userId string, token string, expires time.Duration) *customerror.CustomError {
	key := "user:" + userId
	result := r.client.HSet(context.Background(), key, AccessTokenGroup, token)
	if result.Err() != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(result.Err())
	}
	// set expiration time for the tokens
	err := r.client.Expire(context.Background(), key, expires).Err()
	if err != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(err)
	}
	return nil
}

// SetRefreshToken stores the refresh token for the user
func (r *TokenStore) SetRefreshToken(userId string, token string, expires time.Duration) *customerror.CustomError {
	key := "user:" + userId

	result := r.client.HSet(context.Background(), key, RefreshTokenGroup, token)
	if result.Err() != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(result.Err())
	}
	// set expiration time for the tokens
	err := r.client.Expire(context.Background(), key, expires).Err()
	if err != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(err)
	}
	return nil
}

// GetTokens retrieves the access and refresh tokens for the user
func (r *TokenStore) GetTokens(userId string) (map[string]string, *customerror.CustomError) {
	key := "user:" + userId

	result, err := r.client.HMGet(context.Background(), key, "access_token", "refresh_token").Result()
	if err != nil {
		return map[string]string{}, customerror.ErrorHandler.RedisTokenStoreError(err)
	}

	if result[0] == nil && result[1] == nil {
		return map[string]string{
			"access_token":  "",
			"refresh_token": "",
		}, customerror.ErrorHandler.RedisTokenStoreError(err)
	}

	return map[string]string{
		"access_token":  result[0].(string),
		"refresh_token": result[1].(string),
	}, nil
}

// DeleteToken deletes the access token for the user
func (r *TokenStore) DeleteTokens(userId string) *customerror.CustomError {
	key := "user:" + userId

	result := r.client.Del(context.Background(), key)
	if result.Err() != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(result.Err())
	}
	return nil
}
