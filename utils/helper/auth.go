// Package helper provides auth-related utility functions for the application.
package helper

import (
	"e-commerce/utils/constants"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// CreateJwtWithClaims is used to create a JWT token with claims.
// It takes the data as input and returns the JWT token and a boolean indicating success or failure.
func CreateJwtWithClaims(data any) (string, bool) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "Failed to create auth token", false
	}
	claims[constants.UserJwtClaimKey] = data

	// Set token expiration time (e.g., 1 hour from now)
	expirationTime := time.Now().Add(time.Duration(ExpiryTime) * time.Minute)
	claims["exp"] = expirationTime.Unix()

	jwtToken, err := token.SignedString([]byte(os.Getenv(constants.SecretKey)))
	if err != nil {
		return "Failed to generate auth token", false
	}

	return jwtToken, true
}
