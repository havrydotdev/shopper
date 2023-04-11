package repo

import (
	"github.com/jmoiron/sqlx"
	"shopper"
)

type Authorization interface {
	CreateUser(user shopper.User) (int, error)
	GetUser(email, password string) (shopper.User, error)
	GetUserById(userId int) (shopper.User, error)
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
	GetCompany(companyId int) (shopper.Company, error)
	CreateCompany(company shopper.Company, userId int) (int, error)
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
}

type Repository struct {
	Authorization
	Item
	User
	Rating
	Discount
	Company
	Notification
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		Company:       NewCompanyRepo(db),
		Notification:  NewNotificationRepo(db),
		User:          NewUserRepo(db),
		Discount:      NewDiscountRepo(db),
		Item:          NewItemRepo(db),
		Rating:        NewRatingRepo(db),
		Comment:       NewCommentRepo(db),
	}
}
