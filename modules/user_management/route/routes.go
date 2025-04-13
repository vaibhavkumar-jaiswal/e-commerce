// Package route provides the routing logic for user management operations.
package route

import (
	"e-commerce/middleware/auth"
	"e-commerce/modules/user_management/handler"

	"github.com/gin-gonic/gin"
)

// UserManagementRoutes sets up the routes for user management operations.
func UserManagementRoutes(router *gin.Engine) {
	handler := handler.NewUserHandler()

	router.POST(auth.PublicRoute("/auth/login"), handler.Login)

	router.DELETE("/auth/logout", handler.Logout)
	{
		router.POST(auth.PublicRoute("/user/register"), handler.AddUser)

		router.POST(auth.PublicRoute("/user/verification"), handler.VerifyEmail)

		router.POST(auth.PublicRoute("/user/resend-verification"), handler.ResendVerificationCode)

		router.GET("/user", handler.GetUsers)

		router.GET("/user/:id", handler.GetUserByID)

		router.PUT("/user/:id", handler.UpdateUser)

		router.PATCH("/user/:id", handler.PartialUpdateUser)

		router.DELETE("/user/:id", handler.DeleteUser)
	}

}
