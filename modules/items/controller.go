package items

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tamnk74/todolist-mysql-go/dto"
	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/repository"
)

type ItemController interface {
	listItems(ctx *gin.Context)
	createItem(ctx *gin.Context)
}

type itemController struct {
	itemService ItemService
}

func NewItemController() ItemController {
	itemRepo := repository.NewItemRepository()
	itemService := NewItemService(itemRepo)
	return &itemController{
		itemService: itemService,
	}

}

func (a *itemController) listItems(c *gin.Context) {
	var pagi dto.Pagination
	if err := c.ShouldBindWith(&pagi, binding.Query); err != nil {
		c.Error(err)
		return
	}
	pagi.FillDefault()
	books, _ := a.itemService.ListItems(c.Request.Context(), &pagi)
	pagi.Update()
	c.JSON(200, gin.H{
		"meta": pagi.GetMeta(),
		"data": books,
	})
}

func (a *itemController) createItem(c *gin.Context) {
	var form CreateItem
	if err := c.ShouldBindJSON(&form); err != nil {
		c.Error(err)
		return
	}

	item := models.Item{Name: form.Name}
	newItem, _ := a.itemService.CreateItem(c.Request.Context(), item)
	c.JSON(200, gin.H{"data": newItem})
}
