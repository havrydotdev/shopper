package repo

import "github.com/jmoiron/sqlx"

type DiscountRepo struct {
	db *sqlx.DB
}

func NewDiscountRepo(db *sqlx.DB) *DiscountRepo {
	return &DiscountRepo{
		db: db,
	}
}
