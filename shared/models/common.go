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
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserDetails        UserResponse `json:"user_details"`
	AuthorizationToken string       `json:"token"`
	Expiry             time.Time    `json:"expiry"`
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
