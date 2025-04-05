package requestdata

import (
	"bytes"
	"e-commerce/shared/models"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var RouteStructMap = []models.RouteRegistry{
	{Method: "GET", Path: "/user/:id", Struct: &models.User{}},
	{Method: "POST", Path: "/user", Struct: &models.User{}},
}

func RequestMapping() gin.HandlerFunc {
	return func(context *gin.Context) {
		var targetStruct interface{}
		routeMatched := false

		// Match the request path and method to the registered routes
		for _, route := range RouteStructMap {
			if context.Request.Method == route.Method && context.FullPath() == route.Path {
				targetStruct = route.Struct
				routeMatched = true
				break
			}
		}

		// If no route matched, continue to the next middleware
		if !routeMatched {
			context.Next()
			return
		}

		// Parse the request body into the target struct
		if targetStruct != nil {
			bodyBytes, err := io.ReadAll(context.Request.Body)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
				return
			}

			// Reset the request body for further use
			context.Request.Body = io.NopCloser(io.MultiReader(bytes.NewReader(bodyBytes)))

			// Unmarshal JSON into the target struct
			err = json.Unmarshal(bodyBytes, targetStruct)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
				return
			}
		}

		// Store the mapped struct in the Gin context for access in the handler
		context.Set("requestData", targetStruct)
		context.Next()
	}
}
