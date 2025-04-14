// Package handler provides the HTTP handlers for user management operations.
package handler

import (
	_ "e-commerce/models" // for swagger documentation
	"e-commerce/modules/user/dtos"
	"e-commerce/modules/user/service"
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

// GetUserByID godoc
//
//	@Summary		Get User by ID
//	@Description	Retrieves a user's details by their unique ID.
//	@Tags			Users
//	@Produce		json
//	@Param			id	path		string						true	"User ID"
//	@Success		200	{object}	models.User					"User data fetched successfully"
//	@Failure		400	{object}	shared.BadRequestError		"Invalid ID or user not found"
//	@Failure		500	{object}	shared.InternalServerError	"Internal server error"
//	@Router			/user/{id} [get]
func (handler *Handler) GetUserByID(context *gin.Context) {

	id := context.Param("id")
	if id == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid ID format.")
		return
	}
	user, err := handler.service.GetUserByID(id)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, user)
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

// GetUsers godoc
//
//	@Summary		Get Users with Filters
//	@Description	Returns a list of users with optional filter/query parameters.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			name		query		string						false	"Filter by name"
//	@Param			email		query		string						false	"Filter by email"
//	@Param			is_active	query		bool						false	"Filter by active status"
//	@Success		200			{array}		models.User					"List of users"
//	@Failure		400			{object}	shared.BadRequestError		"Invalid query parameters"
//	@Failure		500			{object}	shared.InternalServerError	"Internal server error"
//	@Router			/users [get]
func (handler *Handler) GetUsers(context *gin.Context) {

	queryParams := &dtos.UserQueryParams{}

	if err := context.ShouldBindQuery(queryParams); err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	users, err := handler.service.GetUsers(queryParams)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, users)
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

// its a generic function used for update and partial update
func handleUserUpdate[T any](
	context *gin.Context,
	updateService func(string, T) (string, error),
) {
	id := context.Param("id")
	if id == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid ID.")
		return
	}

	request, err := helper.BindAndValidate[T](context)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err)
		return
	}

	message, updateErr := updateService(id, request)
	if updateErr != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, updateErr.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, message)
}

// UpdateUser godoc
//
//	@Summary		Update user by ID
//	@Description	Update user details based on the provided ID and request body
//	@Tags			Users
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"User ID"
//	@Param			body	body		dtos.UpdateUserRequest		true	"Update User Request"
//	@Success		200		{object}	dtos.UpdateUserSuccess		"User updated successfully"
//	@Failure		400		{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401		{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500		{object}	shared.InternalServerError	"Unexpected server error"
//	@Router			/users/{id} [put]
func (handler *Handler) UpdateUser(context *gin.Context) {
	handleUserUpdate(
		context,
		handler.service.UpdateUser,
	)
}

// PartialUpdateUser godoc
//
//	@Summary		Partially update user by ID
//	@Description	Partially update user fields using PATCH request
//	@Tags			Users
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"User ID"
//	@Param			body	body		dtos.PatchUserRequest		true	"Partial Update User Request"
//	@Success		200		{object}	dtos.UpdateUserSuccess		"User partially updated successfully"
//	@Failure		400		{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401		{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500		{object}	shared.InternalServerError	"Unexpected server error"
//	@Router			/users/{id} [patch]
func (handler *Handler) PartialUpdateUser(context *gin.Context) {
	handleUserUpdate(
		context,
		handler.service.PartialUpdateUser,
	)
}

// DeleteUser godoc
//
//	@Summary		Delete user by ID
//	@Description	Deletes a user identified by the provided ID
//	@Tags			Users
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string						true	"User ID"
//	@Success		200	{object}	dtos.DeleteUserSuccess		"User deleted successfully"
//	@Failure		400	{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401	{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500	{object}	shared.InternalServerError	"Unexpected server error"
//	@Router			/users/{id} [delete]
func (handler *Handler) DeleteUser(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid ID.")
		return
	}

	message, err := handler.service.DeleteUser(id)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, message)
}
