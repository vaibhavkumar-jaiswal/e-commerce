package helper

import (
	"github.com/gin-gonic/gin"
)

// BindAndValidate binds JSON, query, or path params to the struct and validates it
func BindAndValidate[T any](context *gin.Context) (T, map[string]string) {
	var data T
	if err := context.ShouldBindJSON(&data); err != nil {
		return data, map[string]string{"error": "invalid request format"}
	}

	errorsMap := validateStruct(data)
	if len(errorsMap) > 0 {
		return data, errorsMap
	}
	return data, nil
}
