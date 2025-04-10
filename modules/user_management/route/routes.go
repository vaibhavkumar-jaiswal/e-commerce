package route

import (
	"e-commerce/middleware/auth"
	"e-commerce/modules/user_management/handler"

	"github.com/gin-gonic/gin"
)

func UserManagementRoutes(router *gin.Engine) {
	handler := handler.NewUserHandler()

	router.POST(auth.PublicRoute("/login"), handler.Login)
	{
		user := router.Group("/user")

		user.POST(auth.PublicRoute("/register"), handler.AddUser)

		user.POST(auth.PublicRoute("/verification"), handler.VerifyEmail)

		user.POST(auth.PublicRoute("/resend-verification"), handler.ResendVerificationCode)

		user.GET("", handler.GetUsers)

		user.GET("/:id", handler.GetUserByID)

		user.PUT("/:id", handler.UpdateUser)

		user.PATCH("/:id", handler.PartialUpdateUser)

		user.DELETE("/:id", handler.DeleteUser)
	}

}
