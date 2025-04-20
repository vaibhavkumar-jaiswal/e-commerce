// Package dtos provides Data Transfer Objects (DTOs) for user management operations.
// It includes structures for user registration, login, email verification,
// and other user-related functionalities. These DTOs are used to validate
// incoming requests and format outgoing responses in a consistent manner.
package dtos

import (
	"github.com/google/uuid"
)

// UpdateUserRequest defines fields required when updating an existing user
type UpdateUserRequest struct {
	FirstName string    `json:"first_name" example:"Vaibhavkumar" validate:"required,alpha,min=2,max=20"`
	LastName  string    `json:"last_name" example:"Jaiswal" validate:"required,alpha,min=2,max=20"`
	Email     string    `json:"email" example:"jaiswal.vaibhavkumar45@gmail.com" validate:"required,email"`
	Phone     string    `json:"phone" example:"8888599949" validate:"required,numeric,len=10"`
	RoleID    uuid.UUID `json:"role_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479" validate:"required"`
} //@name UserRequest

// PatchUserRequest allows partial update (PATCH) of user fields
type PatchUserRequest struct {
	FirstName string    `json:"first_name" validate:"omitempty,alpha,min=2,max=50"`
	LastName  string    `json:"last_name" validate:"omitempty,alpha,min=2,max=50"`
	Email     string    `json:"email" validate:"omitempty,email"`
	Phone     string    `json:"phone" validate:"omitempty,numeric,len=10"`
	RoleID    uuid.UUID `json:"role_id" validate:"omitempty"`
} //@name UserRequest

// UserQueryParams defines allowed query parameters for filtering users
type UserQueryParams struct {
	FirstName  *string    `form:"first_name" query:"ILIKE"`
	LastName   *string    `form:"last_name" query:"ILIKE"`
	Email      *string    `form:"email"`
	Phone      *string    `form:"phone"`
	RoleID     *uuid.UUID `form:"role_id"`
	IsVerified bool       `form:"is_verified"`
}

// UpdateUserSuccess represents a successful user update response.
type UpdateUserSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"User upadated successfully."`
} // @name UpdateUserSuccess

// DeleteUserSuccess represents a successful user delete response.
type DeleteUserSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"User deleted successfully."`
} // @name DeleteUserSuccess
