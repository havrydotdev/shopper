package repo

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
	"strings"
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

	query2 := fmt.Sprintf("INSERT INTO %s (name, description, price, amount, keywords, company_id, price_with_discount) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", items)

	var id int
	row := r.db.QueryRow(query2, item.Name, item.Description, item.Price, item.Amount, item.Keywords, company.Id, item.Price)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ItemRepo) GetItems(verified bool) ([]shopper.Item, error) {
	var arr []shopper.Item
	query := fmt.Sprintf("select i.id, i.isverified, i.description, i.company_id, i.amount, i.keywords, i.name, i.rating, i.price, coalesce((i.price / 100) * (100 - (SELECT sum(d.percent) FROM %s INNER JOIN %s "+
		"d on d.id = discounts_items.discount_id WHERE discounts_items.item_id = i.id AND d.relevant > now())), i.price) as price_with_discount from %s i where isverified = $1", discountItems, discounts, items)
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
	query := fmt.Sprintf("select i.id, i.isverified, i.description, i.company_id, i.amount, i.keywords, i.name, i.rating, i.price, coalesce((i.price / 100) * (100 - (SELECT sum(d.percent) FROM %s INNER JOIN %s "+
		"d on d.id = discounts_items.discount_id WHERE discounts_items.item_id = i.id AND d.relevant > now())), i.price) as price_with_discount from %s i where isverified = true AND i.id = $1", discountItems, discounts, items)
	err := r.db.Get(&item, query, id)
	if err != nil {
		return shopper.Item{}, err
	}

	return item, nil
}

func (r *ItemRepo) AddDiscountToItem(id, discountId int) (int, error) {
	var newId int
	query := fmt.Sprintf("INSERT INTO %s (item_id, discount_id) VALUES ($1, $2) RETURNING id", discountItems)
	row := r.db.QueryRow(query, id, discountId)
	err := row.Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (r *ItemRepo) deleteExpired() error {
	query := fmt.Sprintf("DELETE FROM %s WHERE relevant >= now()", discounts)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemRepo) DeleteItem(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s i WHERE company_id = (SELECT company_id FROM %s WHERE id = $1) AND i.id = $2", items, users)
	exec, err := r.db.Exec(query, userId, itemId)
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

func (r *ItemRepo) UpdateItem(userId, itemId int, input shopper.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("amount=$%d", argId))
		args = append(args, *input.Amount)
		argId++
	}

	if input.Keywords != nil {
		setValues = append(setValues, fmt.Sprintf("keywords=$%d", argId))
		args = append(args, *input.Keywords)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s i SET %s FROM %s u WHERE i.company_id = u.company_id AND i.id = $%d AND u.id = $%d", items, setQuery, users, argId, argId+1)

	args = append(args, itemId, userId)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("0 rows affected")
	}

	return nil
}
