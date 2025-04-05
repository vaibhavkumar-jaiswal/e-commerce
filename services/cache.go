// Redis caching
package services

import (
	"e-commerce/database/connections"

	redisCache "github.com/redis/go-redis/v9"
)

func GetRedisClient() *redisCache.Client {
	return connections.GetRedisClient()
}
