package models

import (
	"context"

	"github.com/tamnk74/todolist-mysql-go/schema"
)

// Item type
type Item struct {
	ID     int `gorm:"primary_key"`
	Name   string
	Status int
}

type ItemEntity interface {
	ListItems(ctx context.Context, pagi schema.Pagination) ([]Item, int64, error)
	CreateItem(ctx context.Context, item Item) (Item, error)
}

type ItemRepository interface {
	ListItems(ctx context.Context, pagi schema.Pagination) (res []Item, total int64, err error)
	CreateItem(ctx context.Context, item Item) (Item, error)
}
