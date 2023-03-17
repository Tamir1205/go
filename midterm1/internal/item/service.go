package item

import (
	"context"
	"github.com/Tamir1205/midterm1/internal/storage/items"
)

type service struct {
	itemRepository items.Repository
}

type Service interface {
	FindItem(ctx context.Context, name string) ([]items.Item, error)
}

func NewService(itemRepository items.Repository) Service {
	return &service{itemRepository: itemRepository}
}

func (s *service) FindItem(ctx context.Context, name string) ([]items.Item, error) {
	return s.itemRepository.FindItemsByName(ctx, name)
}
