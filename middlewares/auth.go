package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamnk74/todolist-mysql-go/utils"
)

const identityKey = "id"

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= 7 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		user, ok := utils.VerifyAccessToken(tokenString)
		if ok {
			c.Set("user", user)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
