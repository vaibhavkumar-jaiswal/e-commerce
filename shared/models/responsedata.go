package models

type SuccessResponse[T any] struct {
	Success bool `json:"success" example:"true"`
	Data    T    `json:"data"`
}

type ErrorResponse[T any] struct {
	Success bool `json:"success" example:"false"`
	Error   T    `json:"error" example:"error message"`
}
