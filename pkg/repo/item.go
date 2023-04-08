package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
)

type ItemRepo struct {
	db *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *ItemRepo {
	return &ItemRepo{
		db: db,
	}
}

func (r *ItemRepo) CreateItem(userId int, item shopper.Item) (int, error) {
	var companyId int
	query := fmt.Sprintf("SELECT c.id FROM %s INNER JOIN %s c on c.id = users.company_id WHERE users.id = $1", users, companies)
	err := r.db.Get(&companyId, query, userId)
	if err != nil {
		return 0, err
	}

	query2 := fmt.Sprintf("INSERT INTO %s (name, description, price, amount, keywords, company_id) values ($1, $2, $3, $4, $5, $6) RETURNING id", items)

	var id int
	row := r.db.QueryRow(query2, item.Name, item.Description, item.Price, item.Amount, item.Keywords, companyId)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
