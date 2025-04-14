// Package shared provides swagger models for various API responses.
package shared

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

// LoadDataSuccess represents a successful load data response.
type LoadDataSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Data inserted successfully!"`
} // @name LoadDataSuccess
