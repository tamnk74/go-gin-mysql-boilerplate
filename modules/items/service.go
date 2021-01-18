package items

import (
	"context"

	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/repository"
	"github.com/tamnk74/todolist-mysql-go/schema"
)

type ItemService interface {
	ListItems(ctx context.Context, pagi schema.Pagination) ([]models.Item, int64, error)
	CreateItem(ctx context.Context, item models.Item) (models.Item, error)
}

type itemService struct {
	itemRepo repository.ItemRepository
}

// NewItemService will create new an item object representation of models.Item interface
func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{
		itemRepo: repo,
	}
}

func (a *itemService) ListItems(c context.Context, pagi schema.Pagination) (res []models.Item, total int64, err error) {
	return a.itemRepo.ListItems(c, pagi)
}

func (a *itemService) CreateItem(c context.Context, item models.Item) (res models.Item, err error) {
	return a.itemRepo.CreateItem(c, item)
}
