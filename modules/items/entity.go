package items

import (
	"context"

	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/schema"
)

type ItemEntity struct {
	itemRepo models.ItemRepository
}

// NewItemEntity will create new an item object representation of models.Item interface
func NewItemEntity(repo models.ItemRepository) models.ItemEntity {
	return &ItemEntity{
		itemRepo: repo,
	}
}

func (a *ItemEntity) ListItems(c context.Context, pagi schema.Pagination) (res []models.Item, total int64, err error) {
	return a.itemRepo.ListItems(c, pagi)
}

func (a *ItemEntity) CreateItem(c context.Context, item models.Item) (res models.Item, err error) {
	return a.itemRepo.CreateItem(c, item)
}
