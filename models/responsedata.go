// Package models provides Success and Error response structures
// for API responses. These structures are used to standardize the format of
// API responses across the application.
package models

// SuccessResponse represents a successful API response.
type SuccessResponse[T any] struct {
	Success bool `json:"success" example:"true"`
	Data    T    `json:"data"`
} //@name Success

// ErrorResponse represents an error API response.
type ErrorResponse[T any] struct {
	Success bool `json:"success" example:"false"`
	Error   T    `json:"error"`
} // @name Error
