package handler

import (
	"e-commerce/modules/user_management/service"
	"e-commerce/shared/models"
	"e-commerce/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewUserHandler() *Handler {
	service := service.NewUserService()
	return &Handler{
		service: service,
	}
}

// LoginHandler  godoc
// @Summary      User Login
// @Description  Authenticates a user using email and password. Returns a JWT token on successful login that can be used to authorize future requests.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        loginData  body     models.Login                      true  "User login credentials"
// @Success      200        {object} models.SuccessResponse[LoginResponse] "Authenticated successfully with JWT token"
// @Failure      400        {object} models.BadRequestError                    "Invalid or malformed request body"
// @Failure      401        {object} models.UnauthorizedError                  "Invalid credentials or unauthorized access"
// @Failure      500        {object} models.InternalServerError                "Unexpected server error"
// @Router       /login [post]
func (handler *Handler) Login(context *gin.Context) {
	var loginData models.Login

	err := context.BindJSON(&loginData)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, "Bad request data...!")
		return
	}

	data, err := handler.service.Login(loginData)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, data)
}

// GetUserByID godoc
// @Summary      Get User by ID
// @Description  Retrieves a user's details by their unique ID.
// @Tags         Users
// @Produce      json
// @Param        id   path      string                     true  "User ID"
// @Success      200  {object}  models.User                "User data fetched successfully"
// @Failure      400  {object}  models.BadRequestError     "Invalid ID or user not found"
// @Failure      500  {object}  models.InternalServerError "Internal server error"
// @Router       /user/{id} [get]
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
// @Summary      Verify Email with OTP
// @Description  Verifies a user's email address using an OTP sent to their email.
// @Tags         User Registration
// @Accept       json
// @Produce      json
// @Param        request  body      models.EmailOTPRequest         true  "Email and OTP"
// @Success      200      {object}  models.User                    "User verified successfully"
// @Failure      400      {object}  models.BadRequestError         "Missing or invalid OTP/email"
// @Failure      500      {object}  models.InternalServerError     "Internal server error"
// @Router       /user/verify-email [post]
func (handler *Handler) VerifyEmail(context *gin.Context) {

	request := make(map[string]string)
	err := context.BindJSON(&request)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	otp, ok := request["otp"]
	if !ok || otp == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Please provide OTP for verification.")
		return
	}

	email, ok := request["email"]
	if !ok || email == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Please provide Email for verification.")
		return
	}
	user, err := handler.service.VerifyEmail(email, otp)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, user)
}

// ResendVerificationCode godoc
// @Summary      Resend Verification Code
// @Description  Resends a verification code to the user's email address.
// @Tags         User Registration
// @Accept       json
// @Produce      json
// @Param        request  body      models.ResendEmailRequest      true  "Email for which to resend OTP"
// @Success      200      {object}  models.User                    "OTP sent successfully"
// @Failure      400      {object}  models.BadRequestError         "Missing or invalid email"
// @Failure      500      {object}  models.InternalServerError     "Internal server error"
// @Router       /user/resend-verification [post]
func (handler *Handler) ResendVerificationCode(context *gin.Context) {

	request := make(map[string]string)
	err := context.BindJSON(&request)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	email, ok := request["email"]
	if !ok || email == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Please provide Email for verification.")
		return
	}

	user, err := handler.service.ResendVerificationCode(email)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, user)
}

// GetUsers godoc
// @Summary      Get Users with Filters
// @Description  Returns a list of users with optional filter/query parameters.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        name       query     string  false  "Filter by name"
// @Param        email      query     string  false  "Filter by email"
// @Param        is_active  query     bool    false  "Filter by active status"
// @Success      200        {array}   models.User                    "List of users"
// @Failure      400        {object}  models.BadRequestError         "Invalid query parameters"
// @Failure      500        {object}  models.InternalServerError     "Internal server error"
// @Router       /users [get]
func (handler *Handler) GetUsers(context *gin.Context) {

	queryParams := &models.UserQueryParams{}

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

// RegisterHandler godoc
// @Summary      Register a New User
// @Description  Registers a new user account by accepting valid email and other details. On success, returns the success message. Input validation and uniqueness checks are enforced.
// @Tags         User Registration
// @Accept       json
// @Produce      json
// @Param        userDetails  body     models.UserRequest             true  "User registration payload"
// @Success      200          {object} models.UserRegisterSuccess     "Registration successful"
// @Failure      400          {object} models.BadRequestError         "Invalid input or missing required fields"
// @Failure      401          {object} models.UnauthorizedError       "Unauthorized access attempt"
// @Failure      500          {object} models.InternalServerError     "Unexpected server error"
// @Router       /user/register [post]
func (handler *Handler) AddUser(context *gin.Context) {

	var request models.UserRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid user request data.")
		return
	}

	message, err := handler.service.AddUser(request)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, message)
}

func (handler *Handler) UpdateUser(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid ID.")
		return
	}

	var request models.UpdateUserRequest

	if err := context.ShouldBind(&request); err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid user request data.")
		return
	}

	message, err := handler.service.UpdateUser(id, request)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, message)
}

func (handler *Handler) PartialUpdateUser(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid ID.")
		return
	}

	var request models.PatchUserRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid user request data.")
		return
	}

	message, err := handler.service.PartialUpdateUser(id, request)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, message)
}

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
