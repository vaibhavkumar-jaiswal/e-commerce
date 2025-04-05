package helper

import (
	"context"
	"time"
)

func GetCache(cacheKey string) (string, error) {
	ctx := context.Background()
	data, err := redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return "", err
	}

	return data, nil
}

func DeleteCache(cacheKey ...string) error {
	ctx := context.Background()
	_, err := redisClient.Del(ctx, cacheKey...).Result()
	if err != nil {
		return err
	}

	return nil
}

func SetCache(cacheKey string, value any, expiry time.Duration) (string, error) {
	ctx := context.Background()
	data, err := redisClient.Set(ctx, cacheKey, value, expiry).Result()
	if err != nil {
		return "", err
	}

	return data, nil
}
