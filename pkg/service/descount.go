package service

import "shopper/pkg/repo"

type DiscountService struct {
	repo repo.Discount
}

func NewDiscountService(repo repo.Discount) *DiscountService {
	return &DiscountService{
		repo: repo,
	}
}
