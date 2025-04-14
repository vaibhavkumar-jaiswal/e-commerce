// Package connections manages the initialization, connection, and de-initialization of the Redis client.
package connections

import (
	"context"
	"fmt"

	"e-commerce/shared"

	redisCache "github.com/redis/go-redis/v9"
)

var redisClient *redisCache.Client

// InitRedis initializes the Redis client using the provided Redis connection configuration.
// It verifies the connection by performing a Ping operation.
func InitRedis(redisConnection *shared.RedisConn) error {
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

// DeInitRedis closes the Redis connection and cleans up resources.
func DeInitRedis() error {
	// Close Redis connection
	fmt.Printf("\nClosing redis connection...!")
	if err := redisClient.Close(); err != nil {
		fmt.Printf("\nFailed to close Redis connection: %v", err)
		return err
	}
	return nil
}

// GetRedisClient returns the current Redis client instance.
func GetRedisClient() *redisCache.Client {
	return redisClient
}
