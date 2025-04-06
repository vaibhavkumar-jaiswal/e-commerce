package route

import (
	"e-commerce/modules/user_management/handler"

	"github.com/gin-gonic/gin"
)

func UserManagementRoutes(router *gin.Engine) {
	handler := handler.NewUserHandler()

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
