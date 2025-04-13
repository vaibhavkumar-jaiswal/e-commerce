// Package auth JWT-based authentication middleware for Gin
package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"e-commerce/middleware/ratelimiting"
	"e-commerce/models"

	"e-commerce/utils/constants"
	"e-commerce/utils/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var publicRouteList = map[string]bool{}

// Auth returns a Gin handler function to authenticate requests using JWT.
func Auth() gin.HandlerFunc {
	return authenticate
}

// authenticate is the main function to authenticate requests using JWT.
func authenticate(context *gin.Context) {
	path := context.FullPath()

	if strings.HasPrefix(context.Request.URL.Path, "/api-docs") {
		context.Next()
		return
	}

	if _, ok := publicRouteList[path]; ok {
		context.Next()
		return
	}

	token := context.GetHeader("Authorization")

	if token == "" {
		helper.ResponseWriter(context, http.StatusUnauthorized, "Unauthorized")
		context.Abort()
		return
	}

	tokenParts := strings.SplitN(token, " ", 2)
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		helper.ResponseWriter(context, http.StatusUnauthorized, "Invalid Authorization header format")
		context.Abort()
		return
	}

	isLoggedOut, err := helper.GetCache(constants.LoggedOutKey + ":" + tokenParts[1])
	if err != redis.Nil && err != nil {
		fmt.Println("Error getting cache:", err)
		helper.ResponseWriter(context, http.StatusInternalServerError, "Internal Server Error")
		context.Abort()
		return
	}

	if isLoggedOut == constants.LoggedOutKey {
		helper.ResponseWriter(context, http.StatusUnauthorized, "Unauthorized")
		context.Abort()
		return
	}

	jwtToken, err := jwt.ParseWithClaims(tokenParts[1], jwt.MapClaims{}, func(_ *jwt.Token) (any, error) {
		return []byte(os.Getenv(constants.SecretKey)), nil
	})

	if err != nil {
		helper.ResponseWriter(context, http.StatusUnauthorized, "Unauthorized")
		context.Abort()
		return
	}

	jwtClaims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		helper.ResponseWriter(context, http.StatusUnauthorized, "Unauthorized")
		context.Abort()
		return
	}

	if data, ok := jwtClaims[constants.UserJwtClaimKey].(map[string]any); ok {
		// Convert map to JSON bytes
		jsonData, err := json.Marshal(data)
		if err != nil {
			helper.ResponseWriter(context, http.StatusUnauthorized, "User data conversion failed")
			context.Abort()
			return
		}

		// Decode JSON bytes to User struct
		var userDetails models.User
		err = json.Unmarshal(jsonData, &userDetails)
		if err != nil {
			helper.ResponseWriter(context, http.StatusUnauthorized, "User data conversion failed")
			context.Abort()
			return
		}

		fmt.Printf("Converted to User: %+v\n", userDetails)
		context.Set(constants.UserDataContextKey, userDetails)
		context.Next()
	} else {
		fmt.Println("Conversion to map[string]any failed.")
		helper.ResponseWriter(context, http.StatusUnauthorized, "Conversion to map[string]any failed.")
		context.Abort()
		return
	}
}

// PublicRoute registers a route as public, meaning it doesn't require authentication.
func PublicRoute(route string) string {
	publicRouteList[route] = true
	ratelimiting.PublicRoute(route)
	return route
}
