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
}

type Rating interface {
}

type Discount interface {
}

type Company interface {
	GetCompany(companyId int) (shopper.Company, error)
	CreateCompany(company shopper.Company, userId int) (int, error)
}

type Notification interface {
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
	}
}
