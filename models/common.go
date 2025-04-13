// Package models provides common structs used in the application
package models

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaims represents the claims in the JWT token.
type JWTClaims struct {
	UserDetails any `json:"user_details"`
	jwt.StandardClaims
}
