package repo

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
)

type CommentRepo struct {
	db *sqlx.DB
}

func NewCommentRepo(db *sqlx.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) AddComment(itemId, userId int, input shopper.Comment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (text, item_id, user_id) VALUES ($1, $2, $3) RETURNING id", comments)
	row := r.db.QueryRow(query, input.Text, itemId, userId)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CommentRepo) GetCommentsByItem(id int) ([]shopper.Comment, error) {
	var commentArr []shopper.Comment
	query := fmt.Sprintf("SELECT * FROM %s WHERE item_id = $1", comments)
	err := r.db.Select(&commentArr, query, id)
	if err != nil {
		return nil, err
	}

	return commentArr, nil
}

func (r *CommentRepo) DeleteComment(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND id = $2", comments)
	exec, err := r.db.Exec(query, userId, id)
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

func (r *CommentRepo) UpdateComment(userId, id int, input shopper.UpdateCommentInput) error {
	query := fmt.Sprintf("UPDATE %s SET text = $1 WHERE user_id = $2 AND id = $3", comments)
	exec, err := r.db.Exec(query, input.Text, userId, id)
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
