package models

type SuccessResponse[T any] struct {
	Success bool `json:"success" example:"true"`
	Data    T    `json:"data"`
} //@name Success

type ErrorResponse[T any] struct {
	Success bool `json:"success" example:"false"`
	Error   T    `json:"error"`
} // @name Error
