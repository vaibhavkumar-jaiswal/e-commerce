package responsewrapper

import (
	"e-commerce/utils/constants"

	"github.com/gin-gonic/gin"
)

func GenericResponse() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		// status := context.GetInt(constants.RESPONSE_STATUS_KEY)
		// data, _ := context.Get(constants.RESPONSE_DATA_KEY)
		context.Set(constants.RESPONSE_STATUS_KEY, "")
		context.Set(constants.RESPONSE_DATA_KEY, "")

		// Prepare the generic response
		// var response models.GenericResponse

		// if status >= http.StatusBadRequest {
		// 	response = models.GenericResponse{
		// 		Success: false,
		// 		Error:   data.(string),
		// 	}
		// } else {
		// 	response = models.GenericResponse{
		// 		Success: true,
		// 		Data:    data,
		// 	}
		// }

		// context.JSON(status, response)
	}
}
