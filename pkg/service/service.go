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
}

type Rating interface {
}

type Discount interface {
}

type Company interface {
	CreateCompany(company shopper.Company, userId int) (int, error)
	GetCompanyById(id int) (shopper.Company, error)
}

type Notification interface {
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
	}
}
