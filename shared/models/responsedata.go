package models

type SuccessResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type ErrorResponse[T any] struct {
	Success bool `json:"success"`
	Error   T    `json:"error"`
}
