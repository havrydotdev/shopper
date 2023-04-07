package repo

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
	"strings"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) DeleteUser(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s u WHERE u.id = $1", users)
	res, err := r.db.Exec(query, userId)
	if err != nil {
		return err
	}
	val, err := res.RowsAffected()
	if val == 0 {
		return errors.New("user does not exist")
	}
	return err
}

func (r *UserRepo) GetUserNotifications(userId int) ([]shopper.Notification, error) {
	var notif []shopper.Notification
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", notifications)
	err := r.db.Select(&notif, query, userId)

	return notif, err
}

func (r *UserRepo) GetUserHistory(userId int) ([]shopper.Item, error) {
	var itemList []shopper.Item
	query := fmt.Sprintf("SELECT i.* from %s u INNER JOIN %s ih on u.id = ih.user_id INNER JOIN %s i on i.id = ih.item_id WHERE u.id = $1", users, itemsHistory, items)
	err := r.db.Select(&itemList, query, userId)

	return itemList, err
}

func (r *UserRepo) UpdateUser(userId int, input shopper.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}

	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, *input.Password)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *input.Username)
		argId++
	}

	if input.IsTempBlocked != nil {
		setValues = append(setValues, fmt.Sprintf("istempblocked=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s u SET %s WHERE u.id = $%d", users, setQuery, argId)

	args = append(args, userId)

	res, err := r.db.Exec(query, args...)

	if err != nil {
		return err
	}

	val, err := res.RowsAffected()

	if val == 0 {
		return errors.New("0 rows affected")
	}

	return err
}
