package repo

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
)

type RatingRepo struct {
	db *sqlx.DB
}

func NewRatingRepo(db *sqlx.DB) *RatingRepo {
	return &RatingRepo{
		db: db,
	}
}

func (r *RatingRepo) CreateRate(itemId int, rate shopper.Rate) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (item_id, rate) values ($1, $2) RETURNING id", itemsRatings)
	row := tx.QueryRow(query, itemId, rate.Value)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	upd := fmt.Sprintf("UPDATE %s SET rating = (SELECT AVG(rate) FROM %s WHERE item_id = items.id) WHERE items.id = $1", items, itemsRatings)
	exec, err := tx.Exec(upd, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if affected == 0 {
		tx.Rollback()
		return 0, errors.New("0 rows affected")
	}

	return id, tx.Commit()
}
