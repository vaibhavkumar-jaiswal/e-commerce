package connections

import (
	"context"
	"fmt"

	"e-commerce/shared/models"

	redisCache "github.com/redis/go-redis/v9"
)

var redisClient *redisCache.Client

func InitRedis(redisConnection *models.RedisConn) error {
	redisClient = redisCache.NewClient(&redisCache.Options{
		Addr: redisConnection.Address,
		DB:   redisConnection.DB,
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		fmt.Printf("err connecting to Redis: %#v", err)
		return err
	}

	return nil
}

func DeInitRedis() error {
	// Close Redis connection
	fmt.Printf("\nClosing redis connection...!")
	if err := redisClient.Close(); err != nil {
		fmt.Printf("\nFailed to close Redis connection: %v", err)
		return err
	}
	return nil
}

func GetRedisClient() *redisCache.Client {
	return redisClient
}
