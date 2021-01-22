package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamnk74/todolist-mysql-go/dto"
	"github.com/tamnk74/todolist-mysql-go/utils"
)

const identityKey = "id"

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= 7 {
			c.AbortWithStatusJSON(401, gin.H{
				"errors": []dto.ApiError{
					{
						Code:   "ERR-401",
						Status: http.StatusUnauthorized,
						Title:  http.StatusText(http.StatusUnauthorized),
						Detail: http.StatusText(http.StatusUnauthorized),
					},
				},
			})
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		user, ok := utils.VerifyAccessToken(tokenString)
		if ok {
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatusJSON(401, gin.H{
				"errors": []dto.ApiError{
					{
						Code:   "ERR-401",
						Status: http.StatusUnauthorized,
						Title:  http.StatusText(http.StatusUnauthorized),
						Detail: http.StatusText(http.StatusUnauthorized),
					},
				},
			})
			return
		}
	}
}
