package main

import (
	"e-commerce/modules/user_management/route"

	"github.com/gin-gonic/gin"
)

// registerRoute registers all routes for the application
func registerRoute(router *gin.Engine) {
	route.UserManagementRoutes(router)
}
