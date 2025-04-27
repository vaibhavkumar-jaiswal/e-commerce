package main

import (
	"e-commerce/middleware/authentication"
	authRoute "e-commerce/modules/auth/route"
	productRoute "e-commerce/modules/product/route"
	userRoute "e-commerce/modules/user/route"
	"e-commerce/utils/constants"
	"e-commerce/utils/loaddata"
	"os"

	"github.com/gin-gonic/gin"
)

// registerRoute registers all routes for the application
func registerRoute(router *gin.Engine) {
	v1Router := router.Group(os.Getenv(constants.BasePath))

	authRoute.AuthRoutes(v1Router)

	v1Router.Use(authentication.Auth())
	v1Router.GET(authentication.PublicRoute("/load-data"), loaddata.PreLoadDataHandler)

	userRoute.UserManagementRoutes(v1Router)

	productRoute.ProductRoutes(v1Router)
}
