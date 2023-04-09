package repo

import (
	"errors"
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

type cmp struct {
	Id         int  `db:"id"`
	IsVerified bool `db:"isverified"`
}

func (r *ItemRepo) CreateItem(userId int, item shopper.Item) (int, error) {
	var company cmp
	query := fmt.Sprintf("SELECT c.id, c.isverified FROM %s INNER JOIN %s c on c.id = users.company_id WHERE users.id = $1", users, companies)
	err := r.db.Get(&company, query, userId)
	if err != nil {
		return 0, err
	}

	if !company.IsVerified {
		return 0, errors.New("this company isn`t verified. please wait until your company will be moderated by staff")
	}

	query2 := fmt.Sprintf("INSERT INTO %s (name, description, price, amount, keywords, company_id) values ($1, $2, $3, $4, $5, $6) RETURNING id", items)

	var id int
	row := r.db.QueryRow(query2, item.Name, item.Description, item.Price, item.Amount, item.Keywords, company.Id)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ItemRepo) GetItems(verified bool) ([]shopper.Item, error) {
	var arr []shopper.Item
	query := fmt.Sprintf("SELECT * FROM %s WHERE isverified = $1", items)
	err := r.db.Select(&arr, query, verified)
	if err != nil {
		return nil, err
	}

	return arr, nil
}

func (r *ItemRepo) ModerateItem(id int) error {
	query := fmt.Sprintf("UPDATE %s SET isverified = true WHERE id = $1", items)
	exec, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("0 rows affected")
	}

	return nil
}

func (r *ItemRepo) GetItemById(id int) (shopper.Item, error) {
	var item shopper.Item
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", items)
	err := r.db.Get(&item, query, id)
	if err != nil {
		return shopper.Item{}, err
	}

	return item, nil
}
