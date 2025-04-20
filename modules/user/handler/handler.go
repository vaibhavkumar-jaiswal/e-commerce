// Package handler provides the HTTP handlers for user management operations.
package handler

import (
	_ "e-commerce/models" // for swagger documentation
	"e-commerce/modules/user/dtos"
	"e-commerce/modules/user/service"
	_ "e-commerce/shared" // for swagger documentation
	"e-commerce/utils/helper"
	"net/http"

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

// GetUserByID godoc
//
//	@Summary		Get User by ID
//	@Description	Retrieves a user's details by their unique ID.
//	@Tags			User
//	@Security		BearerAuth
//	@Produce		json
//
//	@Param			user_id	path		string										true	"User ID"
//
//	@Success		200		{object}	shared.SuccessResponse[models.UserResponse]	"User data fetched successfully"
//
//	@Failure		400		{object}	shared.BadRequestError						"Invalid ID or user not found"
//	@Failure		500		{object}	shared.InternalServerError					"Internal server error"
//
//	@Router			/user/{user_id} [get]
func (handler *Handler) GetUserByID(context *gin.Context) {

	userID := context.Param("user_id")
	if userID == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid ID format.")
		return
	}
	user, err := handler.service.GetUserByID(userID)
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
//	@Tags			User
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//
//	@Param			name		query		string											false	"Filter by name"
//	@Param			email		query		string											false	"Filter by email"
//	@Param			is_active	query		bool											false	"Filter by active status"
//
//	@Success		200			{object}	shared.SuccessResponse[[]models.UserResponse]	"List of users"
//
//	@Failure		400			{object}	shared.BadRequestError							"Invalid query parameters"
//	@Failure		500			{object}	shared.InternalServerError						"Internal server error"
//
//	@Router			/user [get]
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

// UpdateUser godoc
//
//	@Summary		Update user by ID
//	@Description	Update user details based on the provided ID and request body
//	@Tags			User
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//
//	@Param			user_id	path		string											true	"User ID"
//	@Param			body	body		dtos.UpdateUserRequest							true	"Update User Request"
//
//	@Success		200		{object}	shared.SuccessResponse[dtos.UpdateUserSuccess]	"User updated successfully"
//
//	@Failure		400		{object}	shared.BadRequestError							"Invalid input or missing required fields"
//	@Failure		401		{object}	shared.UnauthorizedError						"Unauthorized access attempt"
//	@Failure		500		{object}	shared.InternalServerError						"Unexpected server error"
//
//	@Router			/user/{user_id} [put]
func (handler *Handler) UpdateUser(context *gin.Context) {
	helper.HandleUpdateAndPatch(
		"user_id",
		context,
		handler.service.UpdateUser,
	)
}

// PartialUpdateUser godoc
//
//	@Summary		Partially update user by ID
//	@Description	Partially update user fields using PATCH request
//	@Tags			User
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//
//	@Param			user_id	path		string											true	"User ID"
//	@Param			body	body		dtos.PatchUserRequest							true	"Partial Update User Request"
//
//	@Success		200		{object}	shared.SuccessResponse[dtos.UpdateUserSuccess]	"User partially updated successfully"
//
//	@Failure		400		{object}	shared.BadRequestError							"Invalid input or missing required fields"
//	@Failure		401		{object}	shared.UnauthorizedError						"Unauthorized access attempt"
//	@Failure		500		{object}	shared.InternalServerError						"Unexpected server error"
//
//	@Router			/user/{user_id} [patch]
func (handler *Handler) PartialUpdateUser(context *gin.Context) {
	helper.HandleUpdateAndPatch(
		"user_id",
		context,
		handler.service.PartialUpdateUser,
	)
}

// DeleteUser godoc
//
//	@Summary		Delete user by ID
//	@Description	Deletes a user identified by the provided ID
//	@Tags			User
//	@Security		BearerAuth
//	@Produce		json
//
//	@Param			user_id	path		string											true	"User ID"
//
//	@Success		200		{object}	shared.SuccessResponse[dtos.DeleteUserSuccess]	"User deleted successfully"
//
//	@Failure		400		{object}	shared.BadRequestError							"Invalid input or missing required fields"
//	@Failure		401		{object}	shared.UnauthorizedError						"Unauthorized access attempt"
//	@Failure		500		{object}	shared.InternalServerError						"Unexpected server error"
//
//	@Router			/user/{user_id} [delete]
func (handler *Handler) DeleteUser(context *gin.Context) {
	userID := context.Param("user_id")
	if userID == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid ID.")
		return
	}

	message, err := handler.service.DeleteUser(userID)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, message)
}
