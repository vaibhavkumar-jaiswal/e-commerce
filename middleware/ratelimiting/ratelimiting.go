package ratelimiting

import (
	"context"
	"e-commerce/shared/models"
	"e-commerce/utils/constants"
	"e-commerce/utils/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// RateLimiterMiddleware - Limits requests per user (based on IP) in a given time window
func RateLimiter(maxRequests int, timeWindow time.Duration, redisClient *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get user IP
		// userIP := context.ClientIP()
		path := context.FullPath()

		switch path {
		case "/login":
			context.Next()
			return
		case "/user/verification":
			context.Next()
			return
		case "/user/resend-verification":
			context.Next()
			return
		}

		var key string

		if path == "/user/register" {
			key = constants.RATE_LIMIT_PREFIX + context.ClientIP()
		} else {
			userDetails, exists := context.Get(constants.USER_DATA_CONTEXT_KEY)
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
