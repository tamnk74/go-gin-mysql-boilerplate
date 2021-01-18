package items

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.RouterGroup) {
	itemController := NewItemController()
	r.GET("/items", itemController.listItems)
	r.POST("/items", itemController.createItem)
}
