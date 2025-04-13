// Package models provides swagger models for various API responses.
package models

// UnauthorizedError represents an unauthorized error response.
type UnauthorizedError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Unauthorized"`
} // @name UnauthorizedError

// NotFoundError represents a not found error response.
type NotFoundError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Not Found"`
} // @name NotFoundError

// BadRequestError represents a bad request error response.
type BadRequestError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Bad Request"`
} // @name BadRequestError

// InternalServerError represents an internal server error response.
type InternalServerError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Internal Server Error"`
} // @name InternalServerError

// ForbiddenError represents a forbidden error response.
type ForbiddenError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Forbidden"`
} // @name ForbiddenError

// UserRegisterSuccess represents a successful user registration response.
type UserRegisterSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Please verify your Email Address. We have sent an OTP to the Email Address."`
} // @name UserRegisterSuccess

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

// LoadDataSuccess represents a successful load data response.
type LoadDataSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Data inserted successfully!"`
} // @name LoadDataSuccess
