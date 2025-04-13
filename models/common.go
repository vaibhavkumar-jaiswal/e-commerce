// Package models provides common structs used in the application
package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Login represents the login request structure.
type Login struct {
	UserName string `json:"username" validate:"required" example:"john_doe"`
	Password string `json:"password" validate:"required" example:"password123"`
} //@name LoginRequest

// LoginResponse represents the login response structure.
// @swagger:model LoginResponse
type LoginResponse struct {
	UserDetails        UserResponse `json:"user_details"`
	AuthorizationToken string       `json:"token" example:"xxxxxxxxxxxxxxxxxxxxxxxxxxxx_adQssw5c"`
	Expiry             time.Time    `json:"expiry" example:"2025-05-01T12:00:00Z"`
}

// JWTClaims represents the claims in the JWT token.
type JWTClaims struct {
	UserDetails any `json:"user_details"`
	jwt.StandardClaims
}
