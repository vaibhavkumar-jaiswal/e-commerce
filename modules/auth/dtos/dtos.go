// Package dtos provides Data Transfer Objects (DTOs) for user management operations.
// It includes structures for user registration, login, email verification,
// and other user-related functionalities. These DTOs are used to validate
// incoming requests and format outgoing responses in a consistent manner.
package dtos

import (
	"e-commerce/models"
	"time"

	"github.com/google/uuid"
)

// Login-Logout Related Structures

// Login represents the login request structure.
type Login struct {
	UserName string `json:"username" validate:"required" example:"john_doe"`
	Password string `json:"password" validate:"required" example:"password123"`
} //@name LoginRequest

// LoginResponse represents the login response structure.
// @swagger:model LoginResponse
type LoginResponse struct {
	UserDetails        models.UserResponse `json:"user_details"`
	AuthorizationToken string              `json:"token" example:"xxxxxxxxxxxxxxxxxxxxxxxxxxxx_adQssw5c"`
	Expiry             time.Time           `json:"expiry" example:"2025-05-01T12:00:00Z"`
}

// UserRequest defines expected data when creating a new user
type UserRequest struct {
	FirstName string    `json:"first_name" validate:"required,alpha,min=2,max=20"`
	LastName  string    `json:"last_name" validate:"required,alpha,min=2,max=20"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required,numeric,len=10"`
	RoleID    uuid.UUID `json:"role_id" validate:"required"`
} //@name UserRequest

// Email related structures

// EmailOTPRequest represents the request body for sending an OTP to the user's email
type EmailOTPRequest struct {
	Email string `json:"email" validate:"required,email" example:"john.doe@gmail.com"`
	OTP   string `json:"otp" validate:"required" example:"123456"`
} //@name EmailOTPRequest

// ResendEmailRequest represents the request body for resending an OTP to the user's email
type ResendEmailRequest struct {
	Email string `json:"email" validate:"required,email" example:"john.doe@gmail.com"`
} //@name ResendEmailRequest

// EmailVerificationResponse represents the response body for email verification
type EmailVerificationResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Your email has been verified! Login credentials have been sent to your email."`
} //@name EmailVerificationResponse

// UserRegisterSuccess represents a successful user registration response.
type UserRegisterSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Please verify your Email Address. We have sent an OTP to the Email Address."`
} // @name UserRegisterSuccess
