package store

import (
	"context"
	"errors"
	"time"

	"github.com/kokiebisu/rental-storage/service-authentication/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
	"github.com/redis/go-redis/v9"
)

type TokenStore struct {
	client *redis.Client
}

var (
	AccessTokenPrefix  = "access_token_"
	RefreshTokenPrefix = "refresh_token_"
)

// NewTokenStore creates a new token store
func NewTokenStore(client *redis.Client) port.TokenStore {
	return &TokenStore{
		client,
	}
}

// SetAccessToken stores the access token for the user
func (r *TokenStore) SetAccessToken(userID, token string, expires time.Duration) *customerror.CustomError {
	key := AccessTokenPrefix + userID

	result := r.client.Set(context.Background(), key, token, expires)
	if result.Err() != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(result.Err())
	}
	return nil
}

func (r *TokenStore) GetAccessToken(userID string) (string, *customerror.CustomError) {
	key := AccessTokenPrefix + userID

	token, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", customerror.ErrorHandler.RedisTokenStoreError(err)
	}

	return token, nil
}

// DeleteAccessToken deletes the access token for the user
func (r *TokenStore) DeleteAccessToken(userID string) *customerror.CustomError {
	key := AccessTokenPrefix + userID

	result := r.client.Del(context.Background(), key)
	if result.Err() != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(result.Err())
	}
	return nil
}

// SetRefreshToken stores the refresh token for the user
func (r *TokenStore) SetRefreshToken(userID, token string, expires time.Duration) *customerror.CustomError {
	key := RefreshTokenPrefix + userID

	result := r.client.Set(context.Background(), key, token, expires)
	if result.Err() != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(result.Err())
	}
	return nil
}

// GetRefreshToken gets the refresh token for the user
func (r *TokenStore) GetRefreshToken(userID string) (string, *customerror.CustomError) {
	key := RefreshTokenPrefix + userID

	token, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", customerror.ErrorHandler.RedisTokenStoreError(err)
	}

	return token, nil
}

// DeleteRefreshToken deletes the refresh token for the user
func (r *TokenStore) DeleteRefreshToken(userID string) *customerror.CustomError {
	key := RefreshTokenPrefix + userID

	result := r.client.Del(context.Background(), key)
	if result.Err() != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(result.Err())
	}
	return nil
}

// VerifyAccessToken verifies the access token for the user
func (r *TokenStore) VerifyAccessToken(token string, userID string) *customerror.CustomError {
	key := AccessTokenPrefix + userID

	storedToken, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(err)
	}

	if token != storedToken {
		return customerror.ErrorHandler.RedisTokenStoreError(errors.New("invalid token"))
	}

	return nil
}

// VerifyRefreshToken verifies the refresh token for the user
func (r *TokenStore) VerifyRefreshToken(token string, userID string) *customerror.CustomError {
	key := RefreshTokenPrefix + userID

	storedToken, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return customerror.ErrorHandler.RedisTokenStoreError(err)
	}

	if token != storedToken {
		return customerror.ErrorHandler.RedisTokenStoreError(errors.New("invalid token"))
	}

	return nil
}
