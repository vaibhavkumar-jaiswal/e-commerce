// Package route provides the routing logic for auth operations.
package route

import (
	"e-commerce/middleware/auth"
	authHandler "e-commerce/modules/auth/handler"

	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up the routes for user management operations.
func AuthRoutes(router *gin.RouterGroup) {
	handler := authHandler.NewUserHandler()

	router.POST(auth.PublicRoute("/auth/login"), handler.Login)

	router.DELETE("/auth/logout", handler.Logout)

	router.POST(auth.PublicRoute("/auth/register"), handler.AddUser)

	router.POST(auth.PublicRoute("/auth/verification"), handler.VerifyEmail)

	router.POST(auth.PublicRoute("/auth/resend-verification"), handler.ResendVerificationCode)

}
