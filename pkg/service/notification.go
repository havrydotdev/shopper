package service

import (
	"shopper"
	"shopper/pkg/repo"
)

type NotificationService struct {
	repo repo.Notification
}

func (s *NotificationService) CreateNotification(notification shopper.Notification) (int, error) {
	return s.repo.CreateNotification(notification)
}

func NewNotificationService(repo repo.Notification) *NotificationService {
	return &NotificationService{
		repo: repo,
	}
}
