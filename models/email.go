// Package models provides email related models
package models

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
