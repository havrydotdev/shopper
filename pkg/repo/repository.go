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
}

type User interface {
	DeleteUser(userId int) error
	GetUserNotifications(userId int) ([]shopper.Notification, error)
	GetUserHistory(userId int) ([]shopper.Item, error)
	UpdateUser(userId int, input shopper.UpdateUserInput) error
	UpdateUserBalance(userId int, value int) error
}

type Rating interface {
}

type Discount interface {
	CreateDiscount(discount shopper.Discount) (int, error)
}

type Company interface {
	GetCompany(companyId int) (shopper.Company, error)
	CreateCompany(company shopper.Company, userId int) (int, error)
	ModerateCompany(id int) error
	UpdateCompany(userId, companyId int, input shopper.UpdateCompanyInput) error
}

type Notification interface {
	CreateNotification(notification shopper.Notification) (int, error)
}

type Comment interface {
}

type Repository struct {
	Authorization
	Item
	User
	Rating
	Discount
	Company
	Notification
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		Company:       NewCompanyRepo(db),
		Notification:  NewNotificationRepo(db),
		User:          NewUserRepo(db),
		Discount:      NewDiscountRepo(db),
	}
}
