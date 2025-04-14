// Package ratelimiting implements rate limiting middleware for limiting requests
// per user/IP based on a given time window. It uses Redis for storing and tracking request counts.
package ratelimiting

import (
	"context"
	"e-commerce/models"
	"e-commerce/utils/constants"
	"e-commerce/utils/helper"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var publicRouteList = map[string]bool{}
var basePath string

// RateLimiter is a middleware function that limits the number of requests a user can make
// within a specified time window. It uses Redis to track request counts.
func RateLimiter(maxRequests int, timeWindow time.Duration, redisClient *redis.Client) gin.HandlerFunc {
	basePath = os.Getenv(constants.BasePath)
	return func(context *gin.Context) {
		path := context.FullPath()

		var key string

		_, ok := publicRouteList[path]
		if ok {
			key = constants.RateLimitPrefix + context.ClientIP()
		} else {
			userDetails, exists := context.Get(constants.UserDataContextKey)
			if !exists {
				helper.ResponseWriter(context, http.StatusUnauthorized, "Unauthorized")
				context.Abort()
				return
			}

			user, ok := userDetails.(models.User)
			if !ok {
				helper.ResponseWriter(context, http.StatusUnauthorized, "Unauthorized")
				context.Abort()
				return
			}

			// Key for Redis based on user ID
			key = "rate_limit_" + uuid.UUID(user.UserID).String()
		}

		// Get the current count of requests for the user
		count, err := redisClient.Get(ctx, key).Int()
		if err == redis.Nil {
			// No requests made yet, initialize count to 1
			redisClient.Set(ctx, key, 1, timeWindow)
		} else if err != nil {
			// Redis error
			helper.ResponseWriter(context, http.StatusInternalServerError, "Something went wrong, please try again.")
			context.Abort()
			return
		} else if count >= maxRequests {
			// Exceeded rate limit
			helper.ResponseWriter(context, http.StatusTooManyRequests, "Too many requests. Please wait before trying again.")
			context.Abort()
			return
		} else {
			// Increment request count
			redisClient.Incr(ctx, key)
		}

		// Continue to the next middleware/handler
		context.Next()
	}
}

// PublicRoute registers a public route that will use different key to store rate limiting.
// Returns the route string.
func PublicRoute(route string) string {
	publicRouteList[basePath+route] = true
	return route
}
