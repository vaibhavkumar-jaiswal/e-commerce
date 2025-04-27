// Package route provides the routing logic for auth operations.
package route

import (
	"e-commerce/middleware/authentication"
	authHandler "e-commerce/modules/auth/handler"

	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up the routes for user management operations.
func AuthRoutes(router *gin.RouterGroup) {
	handler := authHandler.NewUserHandler()

	router.POST(authentication.PublicRoute("/auth/login"), handler.Login)

	router.DELETE("/auth/logout", handler.Logout)

	router.POST(authentication.PublicRoute("/auth/register"), handler.AddUser)

	router.POST(authentication.PublicRoute("/auth/verification"), handler.VerifyEmail)

	router.POST(authentication.PublicRoute("/auth/resend-verification"), handler.ResendVerificationCode)

}
