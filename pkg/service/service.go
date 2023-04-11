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
	CreateItem(userId int, item shopper.Item) (int, error)
	GetItems(verified bool) ([]shopper.Item, error)
	ModerateItem(id int) error
	GetItemById(id int) (shopper.Item, error)
	AddDiscountToItem(id, discountId int) (int, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input shopper.UpdateItemInput) error
}

type User interface {
	DeleteUser(userId int) error
	GetUserNotifications(userId int) ([]shopper.Notification, error)
	GetUserHistory(userId int) ([]shopper.Item, error)
	UpdateUser(userId int, input shopper.UpdateUserInput) error
	UpdateUserBalance(userId int, value int) error
	ReturnItem(userId, itemId int) error
	BuyItem(userId, itemId int) error
}

type Rating interface {
	CreateRate(itemId int, rate shopper.Rate) (int, error)
}

type Discount interface {
	CreateDiscount(discount shopper.Discount) (int, error)
}

type Company interface {
	CreateCompany(company shopper.Company, userId int) (int, error)
	GetCompanyById(id int) (shopper.Company, error)
	ModerateCompany(id int) error
	UpdateCompany(userId, companyId int, input shopper.UpdateCompanyInput) error
	GetCompanies(verified bool) ([]shopper.Company, error)
}

type Notification interface {
	CreateNotification(notification shopper.Notification) (int, error)
}

type Comment interface {
	AddComment(itemId, userId int, input shopper.Comment) (int, error)
	GetCommentsByItem(id int) ([]shopper.Comment, error)
	DeleteComment(userId, id int) error
	UpdateComment(userId, id int, input shopper.UpdateCommentInput) error
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
		Item:          NewItemService(repos.Item),
		Rating:        NewRatingService(repos.Rating),
		Comment:       NewCommentService(repos.Comment),
	}
}
