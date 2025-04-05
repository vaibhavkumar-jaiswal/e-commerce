// JWT-based authentication
package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"e-commerce/shared/models"

	"e-commerce/utils/constants"
	"e-commerce/utils/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return authenticate
}

func authenticate(context *gin.Context) {
	path := context.FullPath()

	if strings.HasPrefix(context.Request.URL.Path, "/api-docs") {
		context.Next()
		return
	}

	switch path {
	case "/login":
		context.Next()
		return
	case "/user/register":
		context.Next()
		return
	case "/user/verification":
		context.Next()
		return
	case "/user/resend-verification":
		context.Next()
		return
	}

	token := context.GetHeader("Authorization")

	if token == "" {
		helper.ResponseWriter(context, http.StatusUnauthorized, "Unauthorized")
		context.Abort()
		return
	}

	tokenParts := strings.SplitN(token, " ", 2)
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		helper.ResponseWriter(context, http.StatusUnauthorized, "Invalid Authorization header format")
		context.Abort()
		return
	}

	jwtToken, err := jwt.ParseWithClaims(tokenParts[1], jwt.MapClaims{}, func(jwtToken *jwt.Token) (any, error) {
		return []byte(os.Getenv(constants.SECRETE_KEY)), nil
	})

	if err != nil {
		helper.ResponseWriter(context, http.StatusUnauthorized, "Unauthorized")
		context.Abort()
		return
	}

	jwtClaims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		helper.ResponseWriter(context, http.StatusUnauthorized, "Unauthorized")
		context.Abort()
		return
	}

	if data, ok := jwtClaims[constants.USER_JWT_CLAIM_KEY].(map[string]any); ok {
		// Convert map to JSON bytes
		jsonData, err := json.Marshal(data)
		if err != nil {
			helper.ResponseWriter(context, http.StatusUnauthorized, "User data conversion failed")
			context.Abort()
			return
		}

		// Decode JSON bytes to User struct
		var userDetails models.User
		err = json.Unmarshal(jsonData, &userDetails)
		if err != nil {
			helper.ResponseWriter(context, http.StatusUnauthorized, "User data conversion failed")
			context.Abort()
			return
		}

		fmt.Printf("Converted to User: %+v\n", userDetails)
		context.Set(constants.USER_DATA_CONTEXT_KEY, userDetails)
		context.Next()
	} else {
		fmt.Println("Conversion to map[string]any failed.")
		helper.ResponseWriter(context, http.StatusUnauthorized, "Conversion to map[string]any failed.")
		context.Abort()
		return
	}
}
