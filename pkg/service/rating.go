package service

import (
	"shopper"
	"shopper/pkg/repo"
)

type RatingService struct {
	repo repo.Rating
}

func NewRatingService(repo repo.Rating) *RatingService {
	return &RatingService{
		repo: repo,
	}
}

func (s *RatingService) CreateRate(itemId int, rate shopper.Rate) (int, error) {
	return s.repo.CreateRate(itemId, rate)
}
