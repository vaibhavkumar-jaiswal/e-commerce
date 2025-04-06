package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/google/uuid"
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
	UserName string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password123"`
}

type LoginResponse struct {
	UserDetails        UserResponse `json:"user_details"`
	AuthorizationToken string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	Expiry             time.Time    `json:"expiry" example:"2025-05-01T12:00:00Z"`
}

type JWTClaims struct {
	UserDetails any `json:"user_details"`
	jwt.StandardClaims
}

// type baseModel struct {
// 	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"updated_at"`
// 	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`
// 	DeletedAt *time.Time `json:"deleted_at"`
// }
