// Package helper provides utility functions for interacting with Redis cache.
package helper

import (
	"context"
	"time"
)

// GetCache retrieves the value associated with the given cache key from Redis.
// It returns the cached data as a string, or an error if the operation fails.
func GetCache(cacheKey string) (string, error) {
	ctx := context.Background()
	data, err := redisClient.Get(ctx, cacheKey).Result()
	return data, err
}

// DeleteCache removes one or more cache keys from Redis.
// Accepts a variadic list of cache keys and returns an error if the operation fails.
func DeleteCache(cacheKey ...string) error {
	ctx := context.Background()
	_, err := redisClient.Del(ctx, cacheKey...).Result()
	if err != nil {
		return err
	}

	return nil
}

// SetCache sets a key-value pair in Redis with the given expiration time.
// Accepts the cache key, value (as any type), and expiry duration.
// Returns the status of the operation as a string, or an error if the operation fails.
func SetCache(cacheKey string, value any, expiry time.Duration) (string, error) {
	ctx := context.Background()
	data, err := redisClient.Set(ctx, cacheKey, value, expiry).Result()
	if err != nil {
		return "", err
	}

	return data, nil
}
