package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data,omitempty"`
	Error   any  `json:"error,omitempty"`
}

type ResponseWithPagination struct {
	Response
	Pagination
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type Login struct {
	UserName string `json:"username" validate:"required" example:"john_doe"`
	Password string `json:"password" validate:"required" example:"password123"`
} //@name LoginRequest

// @swagger:model LoginResponse
type LoginResponse struct {
	UserDetails        UserResponse `json:"user_details"`
	AuthorizationToken string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	Expiry             time.Time    `json:"expiry" example:"2025-05-01T12:00:00Z"`
}

type JWTClaims struct {
	UserDetails any `json:"user_details"`
	jwt.StandardClaims
}
