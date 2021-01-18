package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/tamnk74/todolist-mysql-go/database"
	models "github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/schema"
)

type ItemRepository interface {
	ListItems(ctx context.Context, pagi schema.Pagination) (res []models.Item, total int64, err error)
	CreateItem(ctx context.Context, item models.Item) (models.Item, error)
}

type itemRepository struct {
	Conn *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewItemRepository() ItemRepository {
	return &itemRepository{database.GetDB()}
}

func (m *itemRepository) ListItems(ctx context.Context, pagi schema.Pagination) (res []models.Item, total int64, err error) {
	var items []models.Item
	var count int64
	m.Conn.Limit(pagi.Limit).Offset(int(pagi.Page-1) * pagi.Limit).Find(&items)
	m.Conn.Model(&models.Item{}).Count(&count)
	return items, count, nil
}

func (m *itemRepository) CreateItem(ctx context.Context, item models.Item) (res models.Item, err error) {
	m.Conn.Create(&item)
	m.Conn.Last(&item)

	return item, nil
}
