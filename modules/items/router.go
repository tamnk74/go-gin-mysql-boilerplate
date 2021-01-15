package items

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/repository"
	"gorm.io/gorm"

	"github.com/tamnk74/todolist-mysql-go/schema"
)

type ItemController struct {
	bookRepo models.ItemEntity
}

func RegisterRouters(r *gin.RouterGroup, db *gorm.DB) {
	itemRepo := repository.NewItemRepository(db)
	itemEntity := NewItemEntity(itemRepo)
	itemController := &ItemController{
		bookRepo: itemEntity,
	}
	r.GET("/items", itemController.listItems)
	r.POST("/items", itemController.createItem)
}

func (a *ItemController) listItems(c *gin.Context) {
	var pagi schema.Pagination
	if err := c.ShouldBindWith(&pagi, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books, total, _ := a.bookRepo.ListItems(c.Request.Context(), pagi)

	c.JSON(200, gin.H{
		"meta": gin.H{
			"total": total,
		},
		"data": books,
	})
}

func (a *ItemController) createItem(c *gin.Context) {
	var form CreateItem
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{Name: form.Name}
	newItem, _ := a.bookRepo.CreateItem(c.Request.Context(), item)
	c.JSON(200, gin.H{"data": newItem})
}
