package item

import (
	"context"
	"github.com/Tamir1205/midterm1/internal/comment"
	"github.com/Tamir1205/midterm1/internal/rating"
	"github.com/Tamir1205/midterm1/internal/storage/items"
)

type service struct {
	itemRepository items.Repository
	ratingService  rating.Service
	commentService comment.Service
}

type Service interface {
	FindItem(ctx context.Context, filter []items.Filter) ([]ItemDto, error)
}

func NewService(itemRepository items.Repository, ratingService rating.Service, commentService comment.Service) Service {
	return &service{
		itemRepository: itemRepository,
		ratingService:  ratingService,
		commentService: commentService,
	}
}

func (s *service) FindItem(ctx context.Context, filter []items.Filter) ([]ItemDto, error) {
	items, err := s.itemRepository.FindItemsByName(ctx, filter)
	if err != nil {
		return nil, err
	}

	dtos := make([]ItemDto, 0)
	for _, v := range items {
		dto := MapItemToDto(v)
		getRating, err := s.ratingService.GetRating(ctx, v.Id)
		if err != nil {
			return nil, err
		}
		dto.Rating = getRating
		comments, err := s.commentService.GetCommentsByItemId(ctx, v.Id)
		if err != nil {
			return nil, err
		}

		dto.Comment = comments
		dtos = append(dtos, dto)
	}

	return dtos, err
}
