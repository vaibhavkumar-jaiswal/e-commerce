package route

import (
	"e-commerce/modules/user_management/handler"
	"e-commerce/modules/user_management/repo"
	"e-commerce/modules/user_management/service"

	"github.com/gin-gonic/gin"
)

func UserManagementRoutes(router *gin.Engine) {
	repo := repo.NewUserRepository()
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	router.POST("/login", handler.LoginHandler)
	{
		user := router.Group("/user")
		user.POST("/register", handler.AddUser)
		user.POST("/verification", handler.VerifyEmail)
		user.POST("/resend-verification", handler.ResendVerificationCode)
		user.GET("", handler.GetUsers)
		user.GET("/:id", handler.GetUserByID)
		// user.PUT("/:id", handler.commonHandler)
		// user.DELETE("/:id", handler.commonHandler)
	}

}
