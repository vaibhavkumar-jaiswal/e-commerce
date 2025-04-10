package models

type EmailOTPRequest struct {
	Email string `json:"email" validate:"required,email" example:"john.doe@gmail.com"`
	OTP   string `json:"otp" validate:"required" example:"123456"`
} //@name EmailOTPRequest

type ResendEmailRequest struct {
	Email string `json:"email" validate:"required,email" example:"john.doe@gmail.com"`
} //@name ResendEmailRequest

type EmailVerificationResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Your email has been successfully verified! We've sent your login credentials to your registered email address. Please check your inbox to proceed."`
} //@name EmailVerificationResponse
