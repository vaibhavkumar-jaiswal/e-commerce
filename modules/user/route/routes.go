// Package route provides the routing logic for user management operations.
package route

import (
	userHandler "e-commerce/modules/user/handler"

	"github.com/gin-gonic/gin"
)

// UserManagementRoutes sets up the routes for user management operations.
func UserManagementRoutes(router *gin.RouterGroup) {
	handler := userHandler.NewUserHandler()
	user := router.Group("/user")
	{
		user.GET("", handler.GetUsers)

		user.GET("/:user_id", handler.GetUserByID)

		user.PUT("/:user_id", handler.UpdateUser)

		user.PATCH("/:user_id", handler.PartialUpdateUser)

		user.DELETE("/:user_id", handler.DeleteUser)
	}

}
