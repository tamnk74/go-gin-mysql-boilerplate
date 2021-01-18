package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tamnk74/todolist-mysql-go/middlewares"
	authModule "github.com/tamnk74/todolist-mysql-go/modules/auth"
	itemModule "github.com/tamnk74/todolist-mysql-go/modules/items"
)

func Init(r *gin.Engine) {
	// Public api
	publicApi := r.Group("/api")
	authModule.RegisterRouters(publicApi)

	// Authenticated api
	authApi := r.Group("/api")
	authApi.Use(middlewares.AuthorizeJWT())
	itemModule.RegisterRouters(authApi)

}
