package service

import (
	"shopper"
	"shopper/pkg/repo"
)

type ItemService struct {
	repo repo.Item
}

func NewItemService(repo repo.Item) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) CreateItem(userId int, item shopper.Item) (int, error) {
	return s.repo.CreateItem(userId, item)
}
