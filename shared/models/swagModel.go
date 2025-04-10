package models

type UnauthorizedError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Unauthorized"`
} // @name UnauthorizedError

type NotFoundError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Not Found"`
} // @name NotFoundError

type BadRequestError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Bad Request"`
} // @name BadRequestError

type InternalServerError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Internal Server Error"`
} // @name InternalServerError

type ForbiddenError struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Forbidden"`
} // @name ForbiddenError

type UserRegisterSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Please verify your Email Address. We have sent an OTP to the Email Address."`
} // @name UserRegisterSuccess

type LoadDataSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Data inserted successfully!"`
} // @name LoadDataSuccess
