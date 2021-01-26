package items

import (
	"context"

	"github.com/tamnk74/todolist-mysql-go/dto"
	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/repository"
)

type ItemService interface {
	ListItems(ctx context.Context, pagi *dto.Pagination) ([]models.Item, error)
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

func (a *itemService) ListItems(c context.Context, pagi *dto.Pagination) (res []models.Item, err error) {
	return a.itemRepo.ListItems(c, pagi)
}

func (a *itemService) CreateItem(c context.Context, item models.Item) (res models.Item, err error) {
	return a.itemRepo.CreateItem(c, item)
}
