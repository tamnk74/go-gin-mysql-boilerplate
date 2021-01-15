package repository

import (
	"context"

	"gorm.io/gorm"

	models "github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/schema"
)

type ItemRepository struct {
	Conn *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewItemRepository(Conn *gorm.DB) models.ItemRepository {
	return &ItemRepository{Conn}
}

func (m *ItemRepository) ListItems(ctx context.Context, pagi schema.Pagination) (res []models.Item, total int64, err error) {
	var items []models.Item
	result := m.Conn.Limit(pagi.Limit).Offset(int(pagi.Page-1) * pagi.Limit).Find(&items)

	return items, result.RowsAffected, nil
}

func (m *ItemRepository) CreateItem(ctx context.Context, item models.Item) (res models.Item, err error) {
	m.Conn.Create(&item)
	m.Conn.Last(&item)

	return item, nil
}
