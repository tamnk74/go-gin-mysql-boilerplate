package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.RouterGroup) {
	authController := NewAuthController()
	r.POST("/login", authController.login)
	r.POST("/register", authController.register)
}
