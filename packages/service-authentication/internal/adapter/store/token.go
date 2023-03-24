package store

import (
	"context"
	"errors"
	"time"

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
func NewTokenStore(client *redis.Client) *TokenStore {
	return &TokenStore{
		client,
	}
}

// SetAccessToken stores the access token for the user
func (r *TokenStore) SetAccessToken(userID, token string, expires time.Duration) error {
	key := AccessTokenPrefix + userID

	return r.client.Set(context.Background(), key, token, expires).Err()
}

func (r *TokenStore) GetAccessToken(userID string) (string, error) {
	key := AccessTokenPrefix + userID

	token, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}

// DeleteAccessToken deletes the access token for the user
func (r *TokenStore) DeleteAccessToken(userID string) error {
	key := AccessTokenPrefix + userID

	return r.client.Del(context.Background(), key).Err()
}

// SetRefreshToken stores the refresh token for the user
func (r *TokenStore) SetRefreshToken(userID, token string, expires time.Duration) error {
	key := RefreshTokenPrefix + userID

	return r.client.Set(context.Background(), key, token, expires).Err()
}

// GetRefreshToken gets the refresh token for the user
func (r *TokenStore) GetRefreshToken(userID string) (string, error) {
	key := RefreshTokenPrefix + userID

	token, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}

// DeleteRefreshToken deletes the refresh token for the user
func (r *TokenStore) DeleteRefreshToken(userID string) error {
	key := RefreshTokenPrefix + userID

	return r.client.Del(context.Background(), key).Err()
}

// VerifyAccessToken verifies the access token for the user
func (r *TokenStore) VerifyAccessToken(token string, userID string) error {
	key := AccessTokenPrefix + userID

	storedToken, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if token != storedToken {
		return errors.New("invalid token")
	}

	return nil
}

// VerifyRefreshToken verifies the refresh token for the user
func (r *TokenStore) VerifyRefreshToken(token string, userID string) error {
	key := RefreshTokenPrefix + userID

	storedToken, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if token != storedToken {
		return errors.New("invalid token")
	}

	return nil
}
