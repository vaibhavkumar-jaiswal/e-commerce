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

// Login represents the login request structure.
type Login struct {
	UserName string `json:"username" validate:"required" example:"vk@gmail.com"`
	Password string `json:"password" validate:"required" example:"123"`
} //@name LoginRequest

// LoginResponse represents the login response structure.
type LoginResponse struct {
	UserDetails        models.UserResponse `json:"user_details"`
	AuthorizationToken string              `json:"token" example:"xxxxxxxxxxxxxxxxxxxxxxxxxxxx_adQssw5c"`
	Expiry             time.Time           `json:"expiry" example:"2025-05-01T12:00:00Z"`
} // @name LoginResponse

// AddUserRequest defines expected data when creating a new user
type AddUserRequest struct {
	FirstName string    `json:"first_name" example:"Vaibhavkumar" binding:"required" validate:"required,alpha,min=2,max=20"`
	LastName  string    `json:"last_name"  example:"Jaiswal"      binding:"required" validate:"required,alpha,min=2,max=20"`
	Email     string    `json:"email"      example:"jaiswal.vaibhavkumar45@gmail.com" binding:"required" validate:"required,email"`
	Phone     string    `json:"phone"      example:"8888599949"   binding:"required" validate:"required,numeric,len=10"`
	RoleID    uuid.UUID `json:"role_id"    example:"f47ac10b-58cc-4372-a567-0e02b2c3d479" binding:"required" validate:"required"`
} //@name UserRequest

// UserRegisterSuccess represents a successful user registration response.
type UserRegisterSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Please verify your Email Address. We have sent an OTP to the Email Address."`
} // @name UserRegisterSuccess

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

// EmailResendResponse represents the response body for email verification
type EmailResendResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"We have sent the OTP to your Email address."`
} //@name EmailResendResponse
