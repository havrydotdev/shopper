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

func (s *ItemService) GetItems(verified bool) ([]shopper.Item, error) {
	return s.repo.GetItems(verified)
}

func (s *ItemService) ModerateItem(id int) error {
	return s.repo.ModerateItem(id)
}

func (s *ItemService) GetItemById(id int) (shopper.Item, error) {
	return s.repo.GetItemById(id)
}
