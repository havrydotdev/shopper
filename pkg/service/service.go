package service

import (
	"shopper"
	"shopper/pkg/repo"
)

type Authorization interface {
	CreateUser(user shopper.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUser(userId int) (shopper.User, error)
}

type Item interface {
}

type User interface {
	DeleteUser(userId int) error
	GetUserNotifications(userId int) ([]shopper.Notification, error)
	GetUserHistory(userId int) ([]shopper.Item, error)
	UpdateUser(userId int, input shopper.UpdateUserInput) error
}

type Rating interface {
}

type Discount interface {
}

type Company interface {
	CreateCompany(company shopper.Company, userId int) (int, error)
	GetCompanyById(id int) (shopper.Company, error)
	ModerateCompany(id int) error
	UpdateCompany(userId, companyId int, input shopper.UpdateCompanyInput) error
}

type Notification interface {
	CreateNotification(notification shopper.Notification) (int, error)
}

type Comment interface {
}

type Service struct {
	Authorization
	Item
	User
	Rating
	Discount
	Company
	Notification
	Comment
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Company:       NewCompanyService(repos.Company),
		Notification:  NewNotificationService(repos.Notification),
		User:          NewUserService(repos.User),
		Discount:      NewDiscountService(repos.Discount),
	}
}
