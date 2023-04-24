package rating

import (
	"context"
	"github.com/Tamir1205/midterm1/internal/storage/rating"
	"math"
)

type service struct {
	ratingRepository rating.Repository
}

type Service interface {
	AddRating(ctx context.Context, createCommentDto CreateRatingDto) error
	GetRating(ctx context.Context, itemId int64) (float64, error)
}

func NewService(ratingRepository rating.Repository) Service {
	return &service{ratingRepository: ratingRepository}
}

func (s *service) AddRating(ctx context.Context, createRatingDto CreateRatingDto) error {
	_, err := s.ratingRepository.Create(ctx, rating.Rating{
		UserId: createRatingDto.UserId,
		ItemId: createRatingDto.ItemId,
		Rating: createRatingDto.Rating,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetRating(ctx context.Context, itemId int64) (float64, error) {
	id, err := s.ratingRepository.GetRatingsByItemId(ctx, itemId)
	if err != nil {
		return 0, err
	}

	if len(id) == 0 {
		return 0, nil
	}

	sum := uint(0)
	for _, v := range id {
		sum += v.Rating
	}

	if sum == 0 {
		return 0, nil
	}

	avg := float64(sum) / float64(len(id))

	return roundFloat(avg, 1), nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
