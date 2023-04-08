package service

import (
	"shopper"
	"shopper/pkg/repo"
)

type DiscountService struct {
	repo repo.Discount
}

func NewDiscountService(repo repo.Discount) *DiscountService {
	return &DiscountService{
		repo: repo,
	}
}

func (s *DiscountService) CreateDiscount(discount shopper.Discount) (int, error) {
	return s.repo.CreateDiscount(discount)
}
