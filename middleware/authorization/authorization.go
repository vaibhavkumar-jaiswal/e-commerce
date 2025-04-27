package authorization

// import (
// 	"e-commerce/database/connections"
// 	"e-commerce/utils/constants"
// 	"e-commerce/utils/helper"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// func Authorize(moduleName string) gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		db := connections.GetDB()

// 		user := helper.GetUserDetails(context)

// 		action := context.Request.Method
// 		if action == constants.Patch {
// 			action = constants.Update
// 		}

// 		roleID := user.RoleID

// 		if err != nil {
// 			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
// 			return
// 		}

// 		if count == 0 {
// 			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
// 			return
// 		}

// 		context.Next()
// 	}
// }
