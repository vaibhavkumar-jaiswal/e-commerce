package middlewares

import (
	"encoding/gob"

	"e-commerce/shared/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	store := memstore.NewStore([]byte("user-expense-app#123"))
	gob.Register(models.UserResponse{})
	gob.Register(models.User{})
	return sessions.Sessions("user-expense-app-session", store)
}
