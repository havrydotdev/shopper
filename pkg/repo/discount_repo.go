package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
)

type DiscountRepo struct {
	db *sqlx.DB
}

func NewDiscountRepo(db *sqlx.DB) *DiscountRepo {
	return &DiscountRepo{
		db: db,
	}
}

func (r *DiscountRepo) CreateDiscount(discount shopper.Discount) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (percent, relevant) values ($1, $2) RETURNING id", discounts)
	row := r.db.QueryRow(query, discount.Percent, discount.Relevant)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
