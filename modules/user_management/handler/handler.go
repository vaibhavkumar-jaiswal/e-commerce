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

// LoginHandler godoc
// @Summary      Login
// @Description  Logs in a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginData  body     models.Login  true  "Login Data"
// @Success      200        {object} models.SuccessResponse[models.LoginResponse]
// @Failure      401        {object} models.ErrorResponse[string]
// @Router       /login [post]
func (handler *Handler) LoginHandler(context *gin.Context) {
	var loginData models.Login

	err := context.BindJSON(&loginData)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, "Bad request data...!")
		return
	}

	data, err := handler.service.GetUserByConditionWithJoin(loginData)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, data)
}

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
// @Summary      Register a new user
// @Description  Registers a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginData  body     models.Login  true  "Login Data"
// @Success      200        {object} models.SuccessResponse[string]
// @Failure      401        {object} models.ErrorResponse[string]
// @Router       /user/register [post]
func (handler *Handler) AddUser(context *gin.Context) {

	var request models.UserRequest

	if err := context.ShouldBind(&request); err != nil {
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

/*

func (handler *Handler) GetUser(context *gin.Context) {
	// var response models.Response

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// response.Success = false
		// response.Error = "Invalid User ID"
		context.JSON(http.StatusBadRequest, "Invalid User ID")
		return
	}

	user, err := handler.service.GetUserByID(uint(id))
	if err != nil {
		// response.Success = false
		// response.Error = err.Error()
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// response.Success = true
	// response.Data = user
	context.JSON(http.StatusOK, user)
}


type userHandler struct {
	userService *userService
}

func newUserHandler(userService *userService) *userHandler {
	return &userHandler{userService: userService}
}

func (handler *userHandler) commonHandler(context *gin.Context) {
	defer helper.CustomRecovery(context)

	fullPath := context.FullPath()
	method := context.Request.Method

	switch {

	case fullPath == "/user/:id" && method == http.MethodGet:
		id_ := context.Param("id")
		id := helper.StringToInt(id_)
		if id == 0 {
			context.JSON(http.StatusBadRequest, models.Response{Success: false, Error: "User data not found for the requested user id..!"})
			return
		}
		userData, err := handler.userService.GetUserByID(id)
		if err != nil {
			fmt.Printf("exception: %#v", err)
			context.JSON(http.StatusBadRequest, models.Response{Success: false, Error: err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": userData})
		return

	case fullPath == "/user" && method == http.MethodGet:
		userData, err := handler.userService.GetUsers()
		if err != nil {
			fmt.Printf("exception: %#v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"data": "Internal Server Error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": userData})
		return

	case fullPath == "/user" && method == http.MethodPost:
		user := models.User{}
		err := context.BindJSON(&user)
		if err != nil {
			context.JSON(http.StatusBadRequest, models.Response{Success: false, Error: "Invalid request data..!"})
			return
		}
		userData, err := handler.userService.AddUser(&user)
		if err != nil {
			fmt.Printf("exception: %#v", err)
			if pgErr, ok := err.(*pq.Error); ok {
				context.JSON(http.StatusBadRequest, models.Response{Success: false, Error: pgErr.Detail})
			} else {
				fmt.Println("Unexpected error type:", err)
				context.JSON(http.StatusBadRequest, models.Response{Success: false, Error: err})
			}
			return
		}
		context.JSON(http.StatusCreated, gin.H{"data": userData})
		return

	case fullPath == "/user/:id" && method == http.MethodPatch:
		id_ := context.Param("id")
		id := helper.StringToInt(id_)
		if id == 0 {
			context.JSON(http.StatusBadRequest, models.Response{Success: false, Error: "User data not found for the requested user id..!"})
			return
		}
		userData, err := handler.userService.GetUserByID(id)
		if err != nil {
			fmt.Printf("exception: %#v", err)
			context.JSON(500, gin.H{"data": "Internal Server Error"})
			return
		}
		context.JSON(200, gin.H{"data": userData})
		return

	case fullPath == "/user/:id" && method == http.MethodDelete:
		id_ := context.Param("id")
		id := helper.StringToInt(id_)
		if id == 0 {
			context.JSON(http.StatusBadRequest, models.Response{Success: false, Error: "User data not found for the requested user id..!"})
			return
		}
		userData, err := handler.userService.GetUserByID(id)
		if err != nil {
			fmt.Printf("exception: %#v", err)
			context.JSON(500, gin.H{"data": "Internal Server Error"})
			return
		}
		context.JSON(200, gin.H{"data": userData})
		return

	default:
		return
	}
}

func (handler *userHandler) loginHandler(context *gin.Context) {
	var response models.Response
	var logindata models.Login

	err := context.BindJSON(&logindata)
	if err != nil {
		response.Success = false
		response.Data = "Bad request data...!"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	userObj, err := handler.userService.LoginUser(logindata)
	if err != nil {
		response.Success = false
		response.Data = err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_details"] = userObj

	// Set token expiration time (e.g., 1 hour from now)
	expirationTime := time.Now().Add(20 * time.Minute)
	claims["exp"] = expirationTime.Unix()

	tokenString, err := token.SignedString([]byte(constants.SECRETE_KEY))
	if err != nil {
		response.Success = false
		response.Data = "Failed to generate auth token"
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	userDetails := userObj.ResponseObj()

	var userData models.LoginResponse
	userData.UserDetails = userDetails
	userData.AuthorizationToken = tokenString
	userData.Expiry = expirationTime

	response.Success = true
	response.Data = userData

	context.JSON(http.StatusOK, response)
}
*/
