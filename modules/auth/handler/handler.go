// Package handler provides the HTTP handlers for auth operations.
package handler

import (
	_ "e-commerce/models" // for swagger documentation
	"e-commerce/modules/auth/dtos"
	"e-commerce/modules/auth/service"
	_ "e-commerce/shared" // for swagger documentation
	"e-commerce/utils/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Handler is the struct that contains the user service
// and is responsible for handling user-related requests.
type Handler struct {
	service *service.Service
}

// NewUserHandler returns the user handler
func NewUserHandler() *Handler {
	service := service.NewUserService()
	return &Handler{
		service: service,
	}
}

// Login  godoc
//
//	@Summary		User Login
//	@Description	Authenticates a user using email and password and returns a JWT token.
//	@Tags			Authentication
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			loginData	body		dtos.Login									true	"User login credentials"
//	@Success		200			{object}	shared.SuccessResponse[dtos.LoginResponse]	"Authenticated successfully with JWT token"
//	@Failure		400			{object}	shared.BadRequestError						"Invalid or malformed request body"
//	@Failure		401			{object}	shared.UnauthorizedError					"Invalid credentials or unauthorized access"
//	@Failure		500			{object}	shared.InternalServerError					"Unexpected server error"
//	@Router			/auth/login [post]
func (handler *Handler) Login(context *gin.Context) {

	loginData, validationErr := helper.BindAndValidate[dtos.Login](context)
	if validationErr != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, validationErr)
		return
	}

	data, err := handler.service.Login(loginData)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, data)
}

// Logout godoc
//
//	@Summary		Logout user
//	@Description	Invalidate the user's access token
//	@Tags			Authentication
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	shared.SuccessResponse[string]	"Successfully logged out"
//	@Failure		400	{object}	shared.BadRequestError			"Invalid or malformed request body"
//	@Failure		401	{object}	shared.UnauthorizedError		"Invalid credentials or unauthorized access"
//	@Failure		500	{object}	shared.InternalServerError		"Unexpected server error"
//	@Router			/logout [post]
func (handler *Handler) Logout(context *gin.Context) {

	tokenParts := strings.SplitN(context.GetHeader("Authorization"), " ", 2)

	data, err := handler.service.Logout(tokenParts[1])
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, data)
}

// VerifyEmail godoc
//
//	@Summary		Verify Email with OTP
//	@Description	Verifies a user's email address using an OTP sent to their email.
//	@Tags			User Registration
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dtos.EmailOTPRequest		true	"Email and OTP"
//	@Success		200		{object}	models.User					"User verified successfully"
//	@Failure		400		{object}	shared.BadRequestError		"Missing or invalid OTP/email"
//	@Failure		500		{object}	shared.InternalServerError	"Internal server error"
//	@Router			/user/verify-email [post]
func (handler *Handler) VerifyEmail(context *gin.Context) {

	request, validationErr := helper.BindAndValidate[dtos.EmailOTPRequest](context)
	if validationErr != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, validationErr)
		return
	}

	user, err := handler.service.VerifyEmail(request.Email, request.OTP)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, user)
}

// ResendVerificationCode godoc
//
//	@Summary		Resend Verification Code
//	@Description	Resends a verification code to the user's email address.
//	@Tags			User Registration
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dtos.ResendEmailRequest		true	"Email for which to resend OTP"
//	@Success		200		{object}	models.User					"OTP sent successfully"
//	@Failure		400		{object}	shared.BadRequestError		"Missing or invalid email"
//	@Failure		500		{object}	shared.InternalServerError	"Internal server error"
//	@Router			/user/resend-verification [post]
func (handler *Handler) ResendVerificationCode(context *gin.Context) {

	request, validationErr := helper.BindAndValidate[dtos.EmailOTPRequest](context)
	if validationErr != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, validationErr)
		return
	}

	user, err := handler.service.ResendVerificationCode(request.Email)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, user)
}

// AddUser godoc
//
//	@Summary		Register a New User
//	@Description	Registers a new user account by accepting valid email and other details.
//	@Tags			User Registration
//	@Accept			json
//	@Produce		json
//	@Param			userDetails	body		dtos.UserRequest			true	"User registration payload"
//	@Success		200			{object}	dtos.UserRegisterSuccess	"Registration successful"
//	@Failure		400			{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401			{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500			{object}	shared.InternalServerError	"Unexpected server error"
//	@Router			/user/register [post]
func (handler *Handler) AddUser(context *gin.Context) {

	request, validationErr := helper.BindAndValidate[dtos.UserRequest](context)
	if validationErr != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, validationErr)
		return
	}

	message, err := handler.service.AddUser(request)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, message)
}
